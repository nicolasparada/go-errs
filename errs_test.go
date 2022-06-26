package errs

import (
	"errors"
	"testing"
)

func TestError_Is(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		target error
		want   bool
	}{
		{
			name:   "unauthenticated",
			err:    UnauthenticatedError("an unauthenticated error"),
			target: ErrUnauthenticated,
			want:   true,
		},
		{
			name:   "not_unauthenticated",
			err:    errors.New("another unauthenticated error"),
			target: ErrUnauthenticated,
			want:   false,
		},
		{
			name:   "invalid_argument",
			err:    InvalidArgumentError("an invalid argument error"),
			target: ErrInvalidArgument,
			want:   true,
		},
		{
			name:   "not_invalid_argument",
			err:    errors.New("another invalid argument error"),
			target: ErrInvalidArgument,
			want:   false,
		},
		{
			name:   "not_found",
			err:    NotFoundError("a not found error"),
			target: ErrNotFound,
			want:   true,
		},
		{
			name:   "not_not_found",
			err:    errors.New("another not found error"),
			target: ErrNotFound,
			want:   false,
		},
		{
			name:   "conflict",
			err:    ConflictError("a conflict error"),
			target: ErrConflict,
			want:   true,
		},
		{
			name:   "not_conflict",
			err:    errors.New("another conflict error"),
			target: ErrConflict,
			want:   false,
		},
		{
			name:   "permission_denied",
			err:    PermissionDeniedError("a permission denied error"),
			target: ErrPermissionDenied,
			want:   true,
		},
		{
			name:   "not_permission_denied",
			err:    errors.New("another permission denied error"),
			target: ErrPermissionDenied,
			want:   false,
		},
		{
			name:   "gone",
			err:    GoneError("a gone error"),
			target: ErrGone,
			want:   true,
		},
		{
			name:   "not_gone",
			err:    errors.New("another gone error"),
			target: ErrGone,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errors.Is(tt.err, tt.target); got != tt.want {
				t.Errorf("errors.Is(%v, %v) = %v, want %v", tt.err, tt.target, got, tt.want)
			}
		})
	}
}
