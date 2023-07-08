package models

import (
	"math/rand"

	"gorm.io/gorm"
  "github.com/graphql-go/graphql"
)

const ALLOWED_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
const ALIAS_LENGTH = 10

type Url struct {
	gorm.Model

	Alias string `gorm:"unique;not null" json:"alias"`
	Url   string `gorm:"not null" json:"url"`
}

func (u Url) GraphQLType() *graphql.Object {
  return graphql.NewObject(graphql.ObjectConfig{
    Name: "Url",
    Fields: graphql.Fields{
      "alias": &graphql.Field{
        Type: graphql.String,
      },
      "url": &graphql.Field{
        Type: graphql.String,
      },
    },
  })
}

func generateAlias() string {
	alias := ""
	for i := 0; i < ALIAS_LENGTH; i++ {
		alias += string(ALLOWED_CHARS[rand.Intn(len(ALLOWED_CHARS))])
	}
	return alias
}

func NewUrl(url string) *Url {
	alias := generateAlias()

	return &Url{
		Alias: alias,
		Url:   url,
	}
}

