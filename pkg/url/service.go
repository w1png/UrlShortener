package url

import "github.com/w1png/urlshortener/storage"

type Url struct {
  Url string `json:"url"`
  Alias string `json:"alias"`
}

type Service interface {
  CreateUrl(url string) (Url, storage.StorageError)
  GetUrl(alias string) (Url, storage.StorageError)
}

