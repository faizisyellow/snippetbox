package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// Create a custom Form struct, which anonymously embeds a url.Values object
// (to hold the form data) and an Errors field to hold any validation errors
// for the form data.
type Form struct {
	url.Values
	Errors errors
}

// Define a New function to initialize a custom Form struct. Notice that
// this takes the form data as the parameter?
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Implement a Required method to check that specific fields in the form
// data are present and not blank. If any fields fail this check, add the
// appropriate message to the form errors.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Implement a MaxLength method to check that a specific field in the form
// contains a maximum number of characters. If the check fails then add the
// appropriate message to the form errors.
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)

	if value == "" {
		return
	}

	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d)", d))
	}
}

// PermittedValues validates that a form field value matches one of the provided permitted values.
// If the field value is empty, validation is skipped. If the value doesn't match any permitted value,
// an error is added to the form's error collection.

// TODO: I dont know if this is even works :)
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}

	// Create a map for O(1) lookup instead of O(n) linear search
	permittedMap := make(map[string]struct{}, len(opts))
	for _, opt := range opts {
		permittedMap[opt] = struct{}{}
	}

	// Check if value exists in the permitted values
	_, exists := permittedMap[value]
	if !exists {
		f.Errors.Add(field, fmt.Sprintf("Must be one of: %s", strings.Join(opts, ", ")))
	}
}

// Implement a Valid method which returns true if there are no errors.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
