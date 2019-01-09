// This package was auto generated.
// DO NOT EDIT BY YOUR HAND!

package errmsg

import "errors"

// StructNameIsNilErr returns the error.
func StructNameIsNilErr() error {
	return errors.New("[GOWRTR-1] struct name must not be empty, but it gets empty")
}

// StructFieldNameIsEmptyErr returns the error.
func StructFieldNameIsEmptyErr() error {
	return errors.New("[GOWRTR-2] field name must not be empty, but it gets empty")
}

// StructFieldTypeIsEmptyErr returns the error.
func StructFieldTypeIsEmptyErr() error {
	return errors.New("[GOWRTR-3] field type must not be empty, but it gets empty")
}

// FuncParameterNameIsEmptyErr returns the error.
func FuncParameterNameIsEmptyErr() error {
	return errors.New("[GOWRTR-4] func parameter name must not be empty, but it gets empty")
}

// LastFuncParameterTypeIsEmptyErr returns the error.
func LastFuncParameterTypeIsEmptyErr() error {
	return errors.New("[GOWRTR-5] the last func parameter type must not be empty, but it gets empty")
}

// FuncNameIsEmptyError returns the error.
func FuncNameIsEmptyError() error {
	return errors.New("[GOWRTR-6] name of func must not be empty, but it gets empty")
}

// InterfaceNameIsEmptyError returns the error.
func InterfaceNameIsEmptyError() error {
	return errors.New("[GOWRTR-7] name of interface must not be empty, but it gets empty")
}

// ErrsList returns the list of errors.
func ErrsList() []string {
	return []string{
		`[GOWRTR-1] struct name must not be empty, but it gets empty`,
		`[GOWRTR-2] field name must not be empty, but it gets empty`,
		`[GOWRTR-3] field type must not be empty, but it gets empty`,
		`[GOWRTR-4] func parameter name must not be empty, but it gets empty`,
		`[GOWRTR-5] the last func parameter type must not be empty, but it gets empty`,
		`[GOWRTR-6] name of func must not be empty, but it gets empty`,
		`[GOWRTR-7] name of interface must not be empty, but it gets empty`,
	}
}
