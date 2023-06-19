package endpoints

type CreateUrlRequest struct {
  Url string `json:"url"`
}

type CreateUrlResponse struct {
  Url string `json:"url"`
  Alias string `json:"alias"`
}

type GetUrlRequest struct {
  Alias string `json:"alias"`
}

type GetUrlResponse struct {
  Url string `json:"url"`
  Alias string `json:"alias"`
}
