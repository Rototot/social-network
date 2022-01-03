package services

import "context"

type RequestResetPasswordParams struct {
	Email string
}

type ResetPasswordParams struct {
	newPassword string
}

type PasswordRecoveryInterface interface {
	ResetPassword(ctx context.Context, payload ResetPasswordParams) error
	RequestResetPassword(ctx context.Context, payload RequestResetPasswordParams) error
}
