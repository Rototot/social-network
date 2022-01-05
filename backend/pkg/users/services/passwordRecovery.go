package services

import "context"

type RequestResetPasswordParams struct {
	Email string
}

type ResetPasswordParams struct {
	NewPassword string
}

type PasswordRecoveryInterface interface {
	ResetPassword(ctx context.Context, payload ResetPasswordParams) error
	RequestResetPassword(ctx context.Context, payload RequestResetPasswordParams) error
}
