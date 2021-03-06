package teststore_test

import (
	"testing"

	"github.com/kapustaprusta/simple_rest_server/internal/app/model"
	"github.com/kapustaprusta/simple_rest_server/internal/app/store"
	"github.com/kapustaprusta/simple_rest_server/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {
	s := teststore.NewStore()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepositoryFindByEmail(t *testing.T) {
	s := teststore.NewStore()
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
