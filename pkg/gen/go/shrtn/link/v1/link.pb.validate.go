// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: shrtn/link/v1/link.proto

package linkv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _link_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Link with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Link) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Hash

	// no validation rules for Url

	// no validation rules for Description

	return nil
}

// LinkValidationError is the validation error returned by Link.Validate if the
// designated constraints aren't met.
type LinkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LinkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LinkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LinkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LinkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LinkValidationError) ErrorName() string { return "LinkValidationError" }

// Error satisfies the builtin error interface
func (e LinkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLink.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LinkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LinkValidationError{}
