package http

import "net/http"
import "social-network/pkg/common"

var errorStatusMap = map[error]int{
	common.ErrNotFound:      http.StatusNotFound,
	common.ErrAlreadyExists: http.StatusBadRequest,
}

type ErrStatusMap map[error]int

func ErrorCodeStatusMapperFunc(err error, extendsStatusMapper ErrStatusMap) int {

	if status, ok := extendsStatusMapper[err]; ok {
		return status
	}

	if status, ok := errorStatusMap[err]; ok {
		return status
	}

	return http.StatusBadRequest
}
