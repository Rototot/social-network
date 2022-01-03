package http

import "net/http"
import "social-network/pkg/common/constants"

var errorStatusMap = map[error]int{
	constants.ErrNotFound:      http.StatusNotFound,
	constants.ErrAlreadyExists: http.StatusBadRequest,
}

func ErrorCodeStatusMapperFunc(err error, extendsStatusMapper map[error]int) int {

	if status, ok := extendsStatusMapper[err]; ok {
		return status
	}

	if status, ok := errorStatusMap[err]; ok {
		return status
	}

	return http.StatusBadRequest
}
