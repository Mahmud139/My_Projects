package main

import (
	"mahmud139/snippetbox/pkg/models"
)

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}