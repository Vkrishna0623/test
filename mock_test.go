package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUserByID(id int) (User, error) {
	args := m.Called(id)
	return args.Get(0).(User), args.Error(1)
}

func TestUserFetcher(t *testing.T) {
	t.Run("Valid User", func(t *testing.T) {
		mockUserService := new(MockUserService)
		mockUserService.On("GetUserByID", 1).Return(User{ID: 1, Name: "Vamsi"}, nil)
		result, err := FetchUserDetails(mockUserService, 1)

		assert.NoError(t, err)
		assert.Equal(t, "User: Vamsi", result)
		mockUserService.AssertExpectations(t)
	})

	t.Run("Invalid User", func(t *testing.T) {
		mockService := new(MockUserService)
		mockService.On("GetUserByID", 2).Return(User{}, errors.New("User Not Found"))
		result, err := FetchUserDetails(mockService, 2)
		assert.Error(t, err)
		assert.Equal(t, "", result)
		mockService.AssertExpectations(t)
	})
}
