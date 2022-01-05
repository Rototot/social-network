package transport

import (
	"net/http"
	commonHttp "social-network/pkg/common/transport/http"
	"social-network/pkg/users"
)

var extendErrorStatusMap = commonHttp.ErrStatusMap{
	users.ErrUserNotFound:       http.StatusNotFound,
	users.ErrPermissionDenied:   http.StatusForbidden,
	users.ErrAuthRequired:       http.StatusUnauthorized,
	users.ErrUserAlreadyLoginIn: http.StatusBadRequest,
}
