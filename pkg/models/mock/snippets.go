package mock

import (
	"time"

	"faizisyellow.com/snippetbox/pkg/models"
)

var mockSnippet = &models.Snippet{
	ID:      3,
	Title:   "Janine freuling",
	Content: "a youtuber from netherlands!",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (s *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 4, nil
}

func (s *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 3:
		return mockSnippet, nil

	default:
		return nil, models.ErrNoRecords

	}
}

func (s *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
