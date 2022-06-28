package httperrs

import (
	"errors"
	"net/http"

	"github.com/nicolasparada/go-errs"
)

// Code returns the HTTP status code of an error
// that was built using the errs package.
func Code(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch {
	case errors.Is(err, errs.ErrUnauthenticated):
		return http.StatusUnauthorized
	case errors.Is(err, errs.ErrInvalidArgument):
		return http.StatusUnprocessableEntity
	case errors.Is(err, errs.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, errs.ErrConflict):
		return http.StatusConflict
	case errors.Is(err, errs.ErrPermissionDenied):
		return http.StatusForbidden
	case errors.Is(err, errs.ErrGone):
		return http.StatusGone
	}

	return http.StatusInternalServerError
}

// IsInternalServerError reports whether err would
// return 500 internal server error status code.
func IsInternalServerError(err error) bool {
	return Code(err) == http.StatusInternalServerError
}
