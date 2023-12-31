package domain

import "errors"

var (
	// ErrNoRecord no record found in database error
	ErrNoRecord = errors.New("models: no matching record found")
	// ErrInvalidCredentials invalid username/password error
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// ErrDuplicateEmail duplicate email error
	ErrDuplicateEmail = errors.New("models: duplicate email")
	// ErrDuplicateAddress duplicate address error
	ErrDuplicateAddress = errors.New("models: duplicate address")
	// ErrPasswordMismatch password mismatch error
	ErrPasswordMismatch = errors.New("models: password mismatch")
	// ErrWeakPassword weak password error
	ErrWeakPassword = errors.New("models: weak password")
	// ErrInvalidEmail invalid email error
	ErrInvalidEmail = errors.New("models: invalid email")
	// ErrInvalidPassword invalid password error
	ErrInvalidPassword = errors.New("models: invalid password")
	// ErrNotFound not found error
	ErrNotFound = errors.New("models: resource not found")
)
