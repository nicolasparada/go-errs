package httperrs

import (
	"errors"
	"net/http"
	"testing"

	"github.com/nicolasparada/go-errs"
)

func TestCode(t *testing.T) {
	tt := []struct {
		name string
		err  error
		want int
	}{
		{
			name: "unathenticated",
			err:  errs.UnauthenticatedError("not logged in"),
			want: http.StatusUnauthorized,
		},
		{
			name: "invalid_argument",
			err:  errs.InvalidArgumentError("invalid title"),
			want: http.StatusUnprocessableEntity,
		},
		{
			name: "not_found",
			err:  errs.NotFoundError("no such post"),
			want: http.StatusNotFound,
		},
		{
			name: "conflict",
			err:  errs.ConflictError("post already exists"),
			want: http.StatusConflict,
		},
		{
			name: "permission_denied",
			err:  errs.PermissionDeniedError("no permission"),
			want: http.StatusForbidden,
		},
		{
			name: "gone",
			err:  errs.GoneError("post has been deleted"),
			want: http.StatusGone,
		},
		{
			name: "internal",
			err:  errors.New("internal error"),
			want: http.StatusInternalServerError,
		},
		{
			name: "ok",
			err:  nil,
			want: http.StatusOK,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := Code(tc.err); got != tc.want {
				t.Errorf("Code(%v) = %d, want %d", tc.err, got, tc.want)
			}
		})
	}
}
