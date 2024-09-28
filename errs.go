// Package errs exposes a common set of error types.
// It is designed so these errors can be declared as constants.
package errs

import "errors"

// Error interface represents an error that was built using this very package.
// Meaning it exposes a Kind() method.
type Error interface {
	error
	Kind() Kind
}

// Kind enum.
type Kind string

const (
	KindUnauthenticated  Kind = "unauthenticated"
	KindInvalidArgument  Kind = "invalid_argument"
	KindNotFound         Kind = "not_found"
	KindConflict         Kind = "conflict"
	KindPermissionDenied Kind = "permission_denied"
	KindGone             Kind = "gone"
)

type UnauthenticatedError string

func (e UnauthenticatedError) Error() string { return string(e) }
func (e UnauthenticatedError) Kind() Kind    { return KindUnauthenticated }
func (e UnauthenticatedError) Is(target error) bool {
	var err Error
	if errors.As(target, &err) {
		return err.Kind() == KindUnauthenticated
	}
	return false
}

type InvalidArgumentError string

func (e InvalidArgumentError) Error() string { return string(e) }
func (e InvalidArgumentError) Kind() Kind    { return KindInvalidArgument }
func (e InvalidArgumentError) Is(target error) bool {
	var err Error
	if errors.As(target, &err) {
		return err.Kind() == KindInvalidArgument
	}
	return false
}

type NotFoundError string

func (e NotFoundError) Error() string { return string(e) }
func (e NotFoundError) Kind() Kind    { return KindNotFound }
func (e NotFoundError) Is(target error) bool {
	var err Error
	if errors.As(target, &err) {
		return err.Kind() == KindNotFound
	}
	return false
}

type ConflictError string

func (e ConflictError) Error() string { return string(e) }
func (e ConflictError) Kind() Kind    { return KindConflict }
func (e ConflictError) Is(target error) bool {
	var err Error
	if errors.As(target, &err) {
		return err.Kind() == KindConflict
	}
	return false
}

type PermissionDeniedError string

func (e PermissionDeniedError) Error() string { return string(e) }
func (e PermissionDeniedError) Kind() Kind    { return KindPermissionDenied }
func (e PermissionDeniedError) Is(target error) bool {
	var err Error
	if errors.As(target, &err) {
		return err.Kind() == KindPermissionDenied
	}
	return false
}

type GoneError string

func (e GoneError) Error() string { return string(e) }
func (e GoneError) Kind() Kind    { return KindGone }
func (e GoneError) Is(target error) bool {
	var err Error
	if errors.As(target, &err) {
		return err.Kind() == KindGone
	}
	return false
}

const (
	Unauthenticated  = UnauthenticatedError("unauthenticated")
	InvalidArgument  = InvalidArgumentError("invalid argument")
	NotFound         = NotFoundError("not found")
	Conflict         = ConflictError("conflict")
	PermissionDenied = PermissionDeniedError("permission denied")
	Gone             = GoneError("gone")
)
