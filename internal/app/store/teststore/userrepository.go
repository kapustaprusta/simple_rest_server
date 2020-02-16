package teststore

import (
	"github.com/kapustaprusta/simple_rest_server/internal/app/model"
	"github.com/kapustaprusta/simple_rest_server/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, isOk := r.users[email]
	if !isOk {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
