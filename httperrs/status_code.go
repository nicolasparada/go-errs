// Package httperrs exposes HTTP utilities
// so you can quickly get an HTTP status code
// out of an error defined in the errs package.
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

	var e errs.Error
	if !errors.As(err, &e) {
		return http.StatusInternalServerError
	}

	switch e.Kind() {
	case errs.KindUnauthenticated:
		return http.StatusUnauthorized
	case errs.KindInvalidArgument:
		return http.StatusUnprocessableEntity
	case errs.KindNotFound:
		return http.StatusNotFound
	case errs.KindConflict:
		return http.StatusConflict
	case errs.KindPermissionDenied:
		return http.StatusForbidden
	case errs.KindGone:
		return http.StatusGone
	default:
		return http.StatusInternalServerError
	}
}

// IsInternalServerError reports whether err would
// return 500 internal server error status code.
func IsInternalServerError(err error) bool {
	return Code(err) == http.StatusInternalServerError
}
