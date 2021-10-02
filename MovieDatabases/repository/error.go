package repository

import "errors"

var (
	NotFound    error = errors.New("not found")
	BadRequest  error = errors.New("bad request")
	ServerError error = errors.New("server error")
)
