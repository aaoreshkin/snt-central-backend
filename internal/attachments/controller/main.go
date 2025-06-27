package controller

import (
	"mime/multipart"
	"net/http"

	"github.com/go-chi/render"
	"github.com/oreshkindev/snt-central-backend/common"
	"github.com/oreshkindev/snt-central-backend/internal/attachments/model"
)

type (
	Controller struct {
		usecase model.Usecase
	}
)

func New(usecase model.Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

func (controller *Controller) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(64 << 20); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}
	defer r.MultipartForm.RemoveAll()

	// Предварительно выделим память
	attachments := make([]model.Attachment, 0, len(r.MultipartForm.File["attachments"]))

	for _, header := range r.MultipartForm.File["attachments"] {

		attachment, err := controller.processCreate(header)
		if err != nil {
			render.Render(w, r, common.ErrInvalidRequest(err))
			return
		}
		attachments = append(attachments, *attachment)
	}

	render.JSON(w, r, attachments)
}

func (controller *Controller) processCreate(header *multipart.FileHeader) (*model.Attachment, error) {
	body, err := header.Open()
	if err != nil {
		return nil, err
	}
	defer body.Close()

	buf := make([]byte, 512)
	if _, err := body.Read(buf); err != nil {
		return nil, err
	}

	if _, err := body.Seek(0, 0); err != nil {
		return nil, err
	}

	return controller.usecase.Create(header, &body)
}

func (controller *Controller) Find(w http.ResponseWriter, r *http.Request) {
	result, err := controller.usecase.Find()
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}
