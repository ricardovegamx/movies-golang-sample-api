package domain

import "errors"

func NotFoundError(msg string) error {
	return errors.New(msg)
}

func BadRequestError(msg string) error {
	return errors.New(msg)
}
