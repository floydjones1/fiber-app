package services

import (
	"testing"

	"github.com/floydjones1/fiber-app/internal/data"
	"github.com/floydjones1/fiber-app/internal/data/mocks"
)

type MockStore struct {
	User mocks.MockUserStore
}

func setupAuthService() (*MockStore, *AuthService) {
	mocks := new(MockStore)

	service := AuthService{
		Store: data.Stores{
			User: &mocks.User,
		},
	}

	return mocks, &service
}

func Test_login(t *testing.T) {
	setupAuthService()
}
