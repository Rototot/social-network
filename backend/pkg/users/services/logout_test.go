package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func factoryLogoutService(t *testing.T) (
	*LogoutService,
	*gomock.Controller,
	*MockSessionStorageInterface,
) {
	ctrl := gomock.NewController(t)
	sessions := NewMockSessionStorageInterface(ctrl)

	service := NewLogoutService(
		sessions,
	)

	return service, ctrl, sessions
}

func TestLogoutWhenExistsSessionIdThenRemoved(t *testing.T) {
	service, controller, sessions := factoryLogoutService(t)
	defer controller.Finish()

	cxt := context.TODO()
	expectedSessionId := SessionId("test_session_id")

	sessions.EXPECT().Remove(expectedSessionId).Return(nil).Times(1)

	err := service.Logout(cxt, expectedSessionId)

	assert.Nil(t, err)
}
