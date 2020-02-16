package store

import (
	"github.com/kapustaprusta/simple_rest_server/internal/app/model"
)

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
