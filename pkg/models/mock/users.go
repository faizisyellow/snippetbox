package mock

import (
	"time"

	"faizisyellow.com/snippetbox/pkg/models"
)

var mockUser = &models.User{
	ID:      1,
	Name:    "lizzy",
	Email:   "lizzymcalpine@gmail.com",
	Created: time.Now(),
}

type userModel struct{}

func (m *userModel) Insert(name, email, password string) error {
	switch email {
	case "lizzymcalpine@gmail.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *userModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "lizzymcalpine@gmail.com":
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

func (m *userModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecords
	}
}
