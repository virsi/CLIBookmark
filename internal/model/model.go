package model

import (
	"errors"
	"net/url"
	"time"
)

type Bookmark struct {
	ID int `json:"id"`
	Name string `json:"name"`
	URL string `json:"url"`
}


func Validate(bm *Bookmark) (bool, error) {
	if bm.Name == "" {
		return false, errors.New("name is empty")
	}
	if bm.URL == "" {
		return false, errors.New("URL is empty")
	}
	_, err := url.ParseRequestURI(bm.URL)
	if err != nil {
		return false, errors.New("URL is not valid")
	}
	return true, nil
}

func GenerateID(bm *Bookmark) {
	bm.ID = int(time.Now().UnixNano())
}
