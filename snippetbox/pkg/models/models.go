package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid Credentials")
	ErrDuplicateEmail = errors.New("models: duplicate email")
)


type Snippet struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}