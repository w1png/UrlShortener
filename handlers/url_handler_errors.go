package handlers

import "fmt"

type UrlHandlerError interface {
  Error() string
}


type InvalidRequestBodyError struct {
  Message string
}

func (e *InvalidRequestBodyError) Error() string {
  return fmt.Sprintf("Invalid request body: %s", e.Message)
}

func NewInvalidRequestBodyError(message string) *InvalidRequestBodyError {
  return &InvalidRequestBodyError{Message: message}
}


type UniqueAliasError struct {
  Message string
}

func (e *UniqueAliasError) Error() string {
  return fmt.Sprintf("Alias already exists: %s", e.Message)
}

func NewUniqueAliasError(message string) *UniqueAliasError {
  return &UniqueAliasError{Message: message}
}

