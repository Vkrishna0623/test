package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserFetch struct {
	mock.Mock
}

func (m MockUserFetch) GetUserByID(id int) (User, error) {
	args := m.Called(id)
	return args.Get(0).(User), args.Error(1)
}

func TestFetchUser(t *testing.T) {
	t.Run("Valid User", func(t *testing.T) {
		mockUserService := new(MockUserFetch)
		mockUserService.On("GetUserByID", 1).Return(User{ID: 1, Name: "John"})

		result, err := FetchUserDetails(mockUserService, 1)

		assert.NoError(t, err)
		assert.Equal(t, "User: John", result)

		mockUserService.AssertExpectations(t)
	})

	t.Run("Invalid User", func(t *testing.T) {
		mockUserService := new(MockUserFetch)
		mockUserService.On("GetUserByID", 2).Return(User{}, errors.New("USer Not Found"))
		result, err := FetchUserDetails(mockUserService, 2)
		assert.Error(t, err)
		assert.Equal(t, "", result)
		mockUserService.AssertExpectations(t)
	})
}
