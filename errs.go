package errs

// All of these errors can be used as `errors.Is` target.
const (
	ErrUnauthenticated  = UnauthenticatedError("unauthenticated")
	ErrInvalidArgument  = InvalidArgumentError("invalid argument")
	ErrNotFound         = NotFoundError("not found")
	ErrConflict         = ConflictError("conflict")
	ErrPermissionDenied = PermissionDeniedError("permission denied")
	ErrGone             = GoneError("gone")
)

// UnauthenticatedError occurs when the caller of an action is not authenticated.
type UnauthenticatedError string

func (e UnauthenticatedError) Error() string        { return string(e) }
func (e UnauthenticatedError) Is(target error) bool { return target == ErrUnauthenticated }

// InvalidArgumentError occurs when an argument is invalid.
type InvalidArgumentError string

func (e InvalidArgumentError) Error() string        { return string(e) }
func (e InvalidArgumentError) Is(target error) bool { return target == ErrInvalidArgument }

// NotFoundError occurs when a resource is not found.
type NotFoundError string

func (e NotFoundError) Error() string        { return string(e) }
func (e NotFoundError) Is(target error) bool { return target == ErrNotFound }

// ConflictError occurs when a resource is in a conflicting state.
type ConflictError string

func (e ConflictError) Error() string        { return string(e) }
func (e ConflictError) Is(target error) bool { return target == ErrConflict }

// PermissionDeniedError occurs when the action is not authorized.
type PermissionDeniedError string

func (e PermissionDeniedError) Error() string        { return string(e) }
func (e PermissionDeniedError) Is(target error) bool { return target == ErrPermissionDenied }

// GoneError occurs when a resource is gone.
type GoneError string

func (e GoneError) Error() string        { return string(e) }
func (e GoneError) Is(target error) bool { return target == ErrGone }
