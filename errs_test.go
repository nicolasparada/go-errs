package errs

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleError() {
	const err = UnauthenticatedError("an unauthenticated error")
	fmt.Println(err.Kind() == KindUnauthenticated)
	// Output: true
}

func TestError(t *testing.T) {
	tt := []struct {
		name string
		err  error
		want Kind
	}{
		{
			name: "const_unauthenticated",
			err:  Unauthenticated,
			want: KindUnauthenticated,
		},
		{
			name: "unauthenticated",
			err:  UnauthenticatedError("an unauthenticated error"),
			want: KindUnauthenticated,
		},
		{
			name: "const_invalid_argument",
			err:  InvalidArgument,
			want: KindInvalidArgument,
		},
		{
			name: "invalid_argument",
			err:  InvalidArgumentError("an invalid argument error"),
			want: KindInvalidArgument,
		},
		{
			name: "const_not_found",
			err:  NotFound,
			want: KindNotFound,
		},
		{
			name: "not_found",
			err:  NotFoundError("a not found error"),
			want: KindNotFound,
		},
		{
			name: "const_conflict",
			err:  Conflict,
			want: KindConflict,
		},
		{
			name: "conflict",
			err:  ConflictError("a conflict error"),
			want: KindConflict,
		},
		{
			name: "const_permission_denied",
			err:  PermissionDenied,
			want: KindPermissionDenied,
		},
		{
			name: "permission_denied",
			err:  PermissionDeniedError("a permission denied error"),
			want: KindPermissionDenied,
		},
		{
			name: "const_gone",
			err:  Gone,
			want: KindGone,
		},
		{
			name: "gone",
			err:  GoneError("a gone error"),
			want: KindGone,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var e Error
			if !errors.As(tc.err, &e) {
				t.Errorf("errors.As(%v, &e) = false, want true", tc.err)
				return
			}

			if got := e.Kind(); got != tc.want {
				t.Errorf("e.Kind() = %v, want %v", got, tc.want)
			}
		})
	}
}
