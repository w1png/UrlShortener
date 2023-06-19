package url

type Url struct {
  Url string `json:"url"`
  Alias string `json:"alias"`
}

type Service interface {
  CreateUrl(url string) (Url, error)
  GetUrl(alias string) (Url, error)
}

