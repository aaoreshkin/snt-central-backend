package usecase

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/oreshkindev/snt-central-backend/common"
	"github.com/oreshkindev/snt-central-backend/internal/attachments/model"
)

type (
	Usecase struct {
		repository  model.Repository
		destination string
	}
)

func New(repository model.Repository) *Usecase {
	return &Usecase{
		repository:  repository,
		destination: os.Getenv("SERVICE_PATH_FILE"),
	}
}

func (usecase *Usecase) Create(header *multipart.FileHeader, body *multipart.File) (*model.Attachment, error) {
	// Получаем расширение и хеш одновременно
	extension, hex, err := usecase.prepareFileInfo(header)
	if err != nil {
		return nil, err
	}

	// Создаем структуру attachment
	file := model.Attachment{
		Name:      strings.TrimSuffix(header.Filename, filepath.Ext(header.Filename)),
		Hex:       hex,
		Size:      uint64(header.Size),
		Extension: extension,
	}
	// Сохраняем в директорию
	if err := usecase.processCreate(hex, extension, body); err != nil {
		return nil, err
	}

	// Сохраняем в репозиторий
	return usecase.repository.Create(&file)
}

func (usecase *Usecase) prepareFileInfo(header *multipart.FileHeader) (string, string, error) {
	extension, err := common.GetExtension(header)
	if err != nil {
		return "", "", err
	}

	hex, err := common.GenerateHex()
	if err != nil {
		return "", "", err
	}

	return extension, hex, nil
}

func (usecase *Usecase) processCreate(hex, extension string, body *multipart.File) error {
	// Создаем директорию если её нет
	if err := os.MkdirAll(usecase.destination, 0755); err != nil {
		return err
	}

	// Формируем полный путь к файлу
	filename := filepath.Join(usecase.destination, hex+"."+extension)

	// Открываем файл для записи
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Копируем содержимое с буфером
	buf := make([]byte, 64*1024)
	if _, err := io.CopyBuffer(f, *body, buf); err != nil {
		// В случае ошибки, пытаемся удалить частично записанный файл
		os.Remove(filename)
		return err
	}

	return nil
}

func (usecase *Usecase) Find() ([]*model.Attachment, error) {
	return usecase.repository.Find()
}
