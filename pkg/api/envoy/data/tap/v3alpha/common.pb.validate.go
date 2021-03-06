// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/data/tap/v3alpha/common.proto

package envoy_data_tap_v3alpha

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

// Validate checks the field values on Body with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Body) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Truncated

	switch m.BodyType.(type) {

	case *Body_AsBytes:
		// no validation rules for AsBytes

	case *Body_AsString:
		// no validation rules for AsString

	}

	return nil
}

// BodyValidationError is the validation error returned by Body.Validate if the
// designated constraints aren't met.
type BodyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BodyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BodyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BodyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BodyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BodyValidationError) ErrorName() string { return "BodyValidationError" }

// Error satisfies the builtin error interface
func (e BodyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBody.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BodyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BodyValidationError{}
