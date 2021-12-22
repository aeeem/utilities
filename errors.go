package utilities

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
	// ErrUnauthorized will throw if the given request-body or params is not valid
	ErrUnauthorized = errors.New("Unauthorize")
	// Errforbiden will throw if user with the given role does not have permission to access the resource
	ErrForbiden = errors.New("The Resource Is Forbidden")
	// ErrDuplicateLogin  will throw if the given request-body or params is not valid
	ErrDuplicateLogin = errors.New("User Already Login, if its not you please change your password.")
	// ErrDuplicateEmail  will throw if the given request-body or params is not valid
	ErrDuplicateEmail = errors.New("User Email Already Exist. Please Change Your email")
	//ErrNeedUpdate will throw if the current apps need to be updated
	ErrNeedUpdate = errors.New("Please update your apps")
)
