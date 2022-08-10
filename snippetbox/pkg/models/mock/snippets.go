package mock

import (
	"mahmud139/snippetbox/pkg/models"
	"time"
)

var mockSnippet = &models.Snippet{
	ID: 1,
	Title: "An old silent pond",
	Content: "an old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title, content, expires string) (int, error){
	return 2, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default: 
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}

func (m *SnippetModel) Delete(id int) error {
	return nil
}