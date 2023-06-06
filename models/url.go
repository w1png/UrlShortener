package models

import (
	"fmt"
	"math/rand"

	"github.com/w1png/ozontest/utils"
	"gorm.io/gorm"
)

const ALLOWED_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
const ALIAS_LENGTH = 10
const MAX_ATTEMPTS = 20

type Url struct {
	gorm.Model

	Alias string `gorm:"unique;not null" json:"alias"`
	Url   string `gorm:"not null" json:"url"`
}

func generateAlias() string {
	alias := ""
	for i := 0; i < ALIAS_LENGTH; i++ {
		alias += string(ALLOWED_CHARS[rand.Intn(len(ALLOWED_CHARS))])
	}
	return alias
}

func generateUniqueAlias(attempt int) (string, error) {
	if attempt > MAX_ATTEMPTS {
		return "", fmt.Errorf("Could not generate unique alias after %d attempts", MAX_ATTEMPTS)
	}

	alias := generateAlias()
	db := utils.DB

	var url Url
	err := db.Where("alias = ?", alias).First(&url).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return alias, nil
		}
		return "", err
	}

	return generateUniqueAlias(attempt + 1)
}

func NewUrl(url string) (*Url, error) {
	alias, err := generateUniqueAlias(0)
	if err != nil {
		return nil, err
	}

	return &Url{
		Alias: alias,
		Url:   url,
	}, nil
}

func (u *Url) Save() error {
	db := utils.DB

	if err := db.Create(&u).Error; err != nil {
		return err
	}

	return nil
}

func GetUrlByAlias(alias string) (*Url, error) {
	db := utils.DB

	var url Url
	err := db.Where("alias = ?", alias).First(&url).Error
	if err != nil {
		return nil, err
	}

	return &url, nil
}
