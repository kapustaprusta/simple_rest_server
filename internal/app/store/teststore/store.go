package teststore

import (
	"github.com/kapustaprusta/simple_rest_server/internal/app/model"
	"github.com/kapustaprusta/simple_rest_server/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// NewStore ...
func NewStore() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
