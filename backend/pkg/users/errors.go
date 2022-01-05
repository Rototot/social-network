package users

import "errors"

var (
	//ErrUserNotFound = errors.New("user not found")
	//ErrPermissionDenied = errors.New("permission denied")
	//ErrUserAlreadyExist     = errors.New("auth: user already exists")
	ErrUserAlreadyLoginIn = errors.New("auth: user already login in")
	ErrUserNotFound       = errors.New("auth: user not found")
	ErrPermissionDenied   = errors.New("auth: permission denied")
	//ErrAuthFailed           = errors.New("auth: authentication failed, invalid user ID or password")
	ErrAuthRequired = errors.New("auth: authentication required")
)
