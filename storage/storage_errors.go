package storage

type StorageError interface {
  Error() string
}


type NotFoundError struct {
  Message string
}

func (e *NotFoundError) Error() string {
  return e.Message
}

func NewNotFoundError(message string) *NotFoundError {
  return &NotFoundError{
    Message: message,
  }
}


type EmptyAliasError struct {
  Message string
}

func (e *EmptyAliasError) Error() string {
  return e.Message
}

func NewEmptyAliasError(message string) *EmptyAliasError {
  return &EmptyAliasError{
    Message: message,
  }
}


type EmptyUrlError struct {
  Message string
}

func (e *EmptyUrlError) Error() string {
  return e.Message
}

func NewEmptyUrlError(message string) *EmptyUrlError {
  return &EmptyUrlError{
    Message: message,
  }
}


type UrlAlreadyExistsError struct {
  Message string
}

func (e *UrlAlreadyExistsError) Error() string {
  return e.Message
}

func NewUrlAlreadyExistsError(message string) *UrlAlreadyExistsError {
  return &UrlAlreadyExistsError{
    Message: message,
  }
}

type StorageLockedError struct {
  Message string
}

func (e *StorageLockedError) Error() string {
  return e.Message
}
  
func NewStorageLockedError(message string) *StorageLockedError {
  return &StorageLockedError{
    Message: message,
  }
}


type UrlIsNilError struct {
  Message string
}

func (e *UrlIsNilError) Error() string {
  return e.Message
}

func NewUrlIsNilError(message string) *UrlIsNilError {
  return &UrlIsNilError{
    Message: message,
  }
}


type DatabaseConnectionError struct {
  Message string
}

func (e *DatabaseConnectionError) Error() string {
  return e.Message
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
  return e.Message
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
  return e.Message
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
  return e.Message
}

func NewDatabaseQueryError(message string) *DatabaseQueryError {
  return &DatabaseQueryError{
    Message: message,
  }
}

