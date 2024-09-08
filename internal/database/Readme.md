## Database

- Accepts the basic context with cancellation
- Returns a new Connection{} structure with our connection

This implementation uses PgxPool because it allows fine tuning of database behaviour and allows for multi-threading.
It doesn't matter which database driver you want to use, it can be Mongo, Arango, Maria and others.

You can create a new directory or several for different drivers inside. Each directory should contain a file with Dial() function,
and in `package database` you can switch between necessary drivers.

### Usage

```go
// Create a context that is cancellable and cancel it on exit.
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

if connection, err = database.New(ctx); err != nil {
	// There is no need to run the application without connecting to the database.
	panic(err)
}
// Close the connection when the program exits.
defer connection.Close()
```
