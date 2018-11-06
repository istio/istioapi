// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/jwt_auth/v2alpha1/config.proto

package v2alpha1

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on HttpUri with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HttpUri) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Uri

	if v, ok := interface{}(m.GetTimeout()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpUriValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.HttpUpstreamType.(type) {

	case *HttpUri_Cluster:
		// no validation rules for Cluster

	}

	return nil
}

// HttpUriValidationError is the validation error returned by HttpUri.Validate
// if the designated constraints aren't met.
type HttpUriValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpUriValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpUriValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpUriValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpUriValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpUriValidationError) ErrorName() string { return "HttpUriValidationError" }

// Error satisfies the builtin error interface
func (e HttpUriValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpUri.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpUriValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpUriValidationError{}

// Validate checks the field values on DataSource with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *DataSource) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Specifier.(type) {

	case *DataSource_Filename:
		// no validation rules for Filename

	case *DataSource_InlineBytes:
		// no validation rules for InlineBytes

	case *DataSource_InlineString:
		// no validation rules for InlineString

	}

	return nil
}

// DataSourceValidationError is the validation error returned by
// DataSource.Validate if the designated constraints aren't met.
type DataSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataSourceValidationError) ErrorName() string { return "DataSourceValidationError" }

// Error satisfies the builtin error interface
func (e DataSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataSourceValidationError{}

// Validate checks the field values on JwtRule with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *JwtRule) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Issuer

	// no validation rules for Forward

	for idx, item := range m.GetFromHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return JwtRuleValidationError{
					field:  fmt.Sprintf("FromHeaders[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for ForwardPayloadHeader

	switch m.JwksSourceSpecifier.(type) {

	case *JwtRule_RemoteJwks:

		if v, ok := interface{}(m.GetRemoteJwks()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return JwtRuleValidationError{
					field:  "RemoteJwks",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *JwtRule_LocalJwks:

		if v, ok := interface{}(m.GetLocalJwks()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return JwtRuleValidationError{
					field:  "LocalJwks",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// JwtRuleValidationError is the validation error returned by JwtRule.Validate
// if the designated constraints aren't met.
type JwtRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JwtRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JwtRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JwtRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JwtRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JwtRuleValidationError) ErrorName() string { return "JwtRuleValidationError" }

// Error satisfies the builtin error interface
func (e JwtRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJwtRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JwtRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JwtRuleValidationError{}

// Validate checks the field values on RemoteJwks with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RemoteJwks) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHttpUri()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return RemoteJwksValidationError{
				field:  "HttpUri",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCacheDuration()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return RemoteJwksValidationError{
				field:  "CacheDuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RemoteJwksValidationError is the validation error returned by
// RemoteJwks.Validate if the designated constraints aren't met.
type RemoteJwksValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoteJwksValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoteJwksValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoteJwksValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoteJwksValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoteJwksValidationError) ErrorName() string { return "RemoteJwksValidationError" }

// Error satisfies the builtin error interface
func (e RemoteJwksValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoteJwks.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoteJwksValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoteJwksValidationError{}

// Validate checks the field values on JwtHeader with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *JwtHeader) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for ValuePrefix

	return nil
}

// JwtHeaderValidationError is the validation error returned by
// JwtHeader.Validate if the designated constraints aren't met.
type JwtHeaderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JwtHeaderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JwtHeaderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JwtHeaderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JwtHeaderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JwtHeaderValidationError) ErrorName() string { return "JwtHeaderValidationError" }

// Error satisfies the builtin error interface
func (e JwtHeaderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJwtHeader.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JwtHeaderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JwtHeaderValidationError{}

// Validate checks the field values on JwtAuthentication with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *JwtAuthentication) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return JwtAuthenticationValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for AllowMissingOrFailed

	return nil
}

// JwtAuthenticationValidationError is the validation error returned by
// JwtAuthentication.Validate if the designated constraints aren't met.
type JwtAuthenticationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JwtAuthenticationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JwtAuthenticationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JwtAuthenticationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JwtAuthenticationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JwtAuthenticationValidationError) ErrorName() string {
	return "JwtAuthenticationValidationError"
}

// Error satisfies the builtin error interface
func (e JwtAuthenticationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJwtAuthentication.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JwtAuthenticationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JwtAuthenticationValidationError{}
