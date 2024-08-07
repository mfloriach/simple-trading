// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: authentication.proto

package grpc_todo

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on SignInRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignInRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignInRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignInRequestMultiError, or
// nil if none found.
func (m *SignInRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignInRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	if len(errors) > 0 {
		return SignInRequestMultiError(errors)
	}

	return nil
}

// SignInRequestMultiError is an error wrapping multiple validation errors
// returned by SignInRequest.ValidateAll() if the designated constraints
// aren't met.
type SignInRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignInRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignInRequestMultiError) AllErrors() []error { return m }

// SignInRequestValidationError is the validation error returned by
// SignInRequest.Validate if the designated constraints aren't met.
type SignInRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInRequestValidationError) ErrorName() string { return "SignInRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignInRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInRequestValidationError{}

// Validate checks the field values on SignInReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignInReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignInReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignInReplyMultiError, or
// nil if none found.
func (m *SignInReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SignInReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return SignInReplyMultiError(errors)
	}

	return nil
}

// SignInReplyMultiError is an error wrapping multiple validation errors
// returned by SignInReply.ValidateAll() if the designated constraints aren't met.
type SignInReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignInReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignInReplyMultiError) AllErrors() []error { return m }

// SignInReplyValidationError is the validation error returned by
// SignInReply.Validate if the designated constraints aren't met.
type SignInReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInReplyValidationError) ErrorName() string { return "SignInReplyValidationError" }

// Error satisfies the builtin error interface
func (e SignInReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInReplyValidationError{}
