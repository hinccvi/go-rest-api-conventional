// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/auth.proto

package proto

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

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRequestMultiError, or
// nil if none found.
func (m *LoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	if len(errors) > 0 {
		return LoginRequestMultiError(errors)
	}

	return nil
}

// LoginRequestMultiError is an error wrapping multiple validation errors
// returned by LoginRequest.ValidateAll() if the designated constraints aren't met.
type LoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRequestMultiError) AllErrors() []error { return m }

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

// Validate checks the field values on LoginReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReplyMultiError, or
// nil if none found.
func (m *LoginReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return LoginReplyMultiError(errors)
	}

	return nil
}

// LoginReplyMultiError is an error wrapping multiple validation errors
// returned by LoginReply.ValidateAll() if the designated constraints aren't met.
type LoginReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReplyMultiError) AllErrors() []error { return m }

// LoginReplyValidationError is the validation error returned by
// LoginReply.Validate if the designated constraints aren't met.
type LoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReplyValidationError) ErrorName() string { return "LoginReplyValidationError" }

// Error satisfies the builtin error interface
func (e LoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReplyValidationError{}

// Validate checks the field values on RefreshRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RefreshRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefreshRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RefreshRequestMultiError,
// or nil if none found.
func (m *RefreshRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RefreshRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return RefreshRequestMultiError(errors)
	}

	return nil
}

// RefreshRequestMultiError is an error wrapping multiple validation errors
// returned by RefreshRequest.ValidateAll() if the designated constraints
// aren't met.
type RefreshRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefreshRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefreshRequestMultiError) AllErrors() []error { return m }

// RefreshRequestValidationError is the validation error returned by
// RefreshRequest.Validate if the designated constraints aren't met.
type RefreshRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshRequestValidationError) ErrorName() string { return "RefreshRequestValidationError" }

// Error satisfies the builtin error interface
func (e RefreshRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshRequestValidationError{}

// Validate checks the field values on RefreshReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RefreshReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefreshReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RefreshReplyMultiError, or
// nil if none found.
func (m *RefreshReply) ValidateAll() error {
	return m.validate(true)
}

func (m *RefreshReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	if len(errors) > 0 {
		return RefreshReplyMultiError(errors)
	}

	return nil
}

// RefreshReplyMultiError is an error wrapping multiple validation errors
// returned by RefreshReply.ValidateAll() if the designated constraints aren't met.
type RefreshReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefreshReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefreshReplyMultiError) AllErrors() []error { return m }

// RefreshReplyValidationError is the validation error returned by
// RefreshReply.Validate if the designated constraints aren't met.
type RefreshReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshReplyValidationError) ErrorName() string { return "RefreshReplyValidationError" }

// Error satisfies the builtin error interface
func (e RefreshReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshReplyValidationError{}
