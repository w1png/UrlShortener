package utils

import "fmt"

type NotFoundError struct {
  Message string
}

func (e *NotFoundError) Error() string {
  return fmt.Sprintf("Object not found: %s", e.Message)
}

func NewNotFoundError(message string) *NotFoundError {
  return &NotFoundError{
    Message: message,
  }
}


type EmptyAliasError struct {}

func (e *EmptyAliasError) Error() string {
  return "Alias is empty"
}

func NewEmptyAliasError() *EmptyAliasError {
  return &EmptyAliasError{}
}


type EmptyUrlError struct {}

func (e *EmptyUrlError) Error() string {
  return "Url is empty"
}

func NewEmptyUrlError() *EmptyUrlError {
  return &EmptyUrlError{}
}


type UrlAlreadyExistsError struct {
  Message string
}

func (e *UrlAlreadyExistsError) Error() string {
  return fmt.Sprintf("Url already exists: %s", e.Message)
}

func NewUrlAlreadyExistsError(message string) *UrlAlreadyExistsError {
  return &UrlAlreadyExistsError{
    Message: message,
  }
}

type StorageLockedError struct {}

func (e *StorageLockedError) Error() string {
  return "Storage locked"
}
  
func NewStorageLockedError() *StorageLockedError {
  return &StorageLockedError{}
}


type UrlIsNilError struct {}

func (e *UrlIsNilError) Error() string {
  return "Url is nil"
}

func NewUrlIsNilError() *UrlIsNilError {
  return &UrlIsNilError{}
}


type DatabaseConnectionError struct {
  Message string
}

func (e *DatabaseConnectionError) Error() string {
  return fmt.Sprintf("Database connection error: %s", e.Message)
}

func NewDatabaseConnectionError(message string) *DatabaseConnectionError {
  return &DatabaseConnectionError{
    Message: message,
  }
}


type EnvironmentVariableError struct {
  Message string
}

func (e *EnvironmentVariableError) Error() string {
  return fmt.Sprintf("Environment variable error: %s", e.Message)
}

func NewEnvironmentVariableError(message string) *EnvironmentVariableError {
  return &EnvironmentVariableError{
    Message: message,
  }
}


type DatabaseError struct {
  Message string
}

func (e *DatabaseError) Error() string {
  return fmt.Sprintf("Database error: %s", e.Message)
}

func NewDatabaseError(message string) *DatabaseError {
  return &DatabaseError{
    Message: message,
  }
}


type DatabaseQueryError struct {
  Message string
}

func (e *DatabaseQueryError) Error() string {
  return fmt.Sprintf("Database query error: %s", e.Message)
}

func NewDatabaseQueryError(message string) *DatabaseQueryError {
  return &DatabaseQueryError{
    Message: message,
  }
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

