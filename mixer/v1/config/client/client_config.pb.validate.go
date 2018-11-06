// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: mixer/v1/config/client/client_config.proto

package client

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

// Validate checks the field values on NetworkFailPolicy with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *NetworkFailPolicy) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Policy

	return nil
}

// NetworkFailPolicyValidationError is the validation error returned by
// NetworkFailPolicy.Validate if the designated constraints aren't met.
type NetworkFailPolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NetworkFailPolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NetworkFailPolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NetworkFailPolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NetworkFailPolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NetworkFailPolicyValidationError) ErrorName() string {
	return "NetworkFailPolicyValidationError"
}

// Error satisfies the builtin error interface
func (e NetworkFailPolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNetworkFailPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NetworkFailPolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NetworkFailPolicyValidationError{}

// Validate checks the field values on ServiceConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ServiceConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DisableCheckCalls

	// no validation rules for DisableReportCalls

	if v, ok := interface{}(m.GetMixerAttributes()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return ServiceConfigValidationError{
				field:  "MixerAttributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetHttpApiSpec() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return ServiceConfigValidationError{
					field:  fmt.Sprintf("HttpApiSpec[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetQuotaSpec() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return ServiceConfigValidationError{
					field:  fmt.Sprintf("QuotaSpec[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetNetworkFailPolicy()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return ServiceConfigValidationError{
				field:  "NetworkFailPolicy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetForwardAttributes()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return ServiceConfigValidationError{
				field:  "ForwardAttributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ServiceConfigValidationError is the validation error returned by
// ServiceConfig.Validate if the designated constraints aren't met.
type ServiceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceConfigValidationError) ErrorName() string { return "ServiceConfigValidationError" }

// Error satisfies the builtin error interface
func (e ServiceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceConfigValidationError{}

// Validate checks the field values on TransportConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TransportConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DisableCheckCache

	// no validation rules for DisableQuotaCache

	// no validation rules for DisableReportBatch

	if v, ok := interface{}(m.GetNetworkFailPolicy()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return TransportConfigValidationError{
				field:  "NetworkFailPolicy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetStatsUpdateInterval()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return TransportConfigValidationError{
				field:  "StatsUpdateInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for CheckCluster

	// no validation rules for ReportCluster

	if v, ok := interface{}(m.GetAttributesForMixerProxy()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return TransportConfigValidationError{
				field:  "AttributesForMixerProxy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TransportConfigValidationError is the validation error returned by
// TransportConfig.Validate if the designated constraints aren't met.
type TransportConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransportConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransportConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransportConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransportConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransportConfigValidationError) ErrorName() string { return "TransportConfigValidationError" }

// Error satisfies the builtin error interface
func (e TransportConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransportConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransportConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransportConfigValidationError{}

// Validate checks the field values on HttpClientConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HttpClientConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTransport()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpClientConfigValidationError{
				field:  "Transport",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ServiceConfigs

	// no validation rules for DefaultDestinationService

	if v, ok := interface{}(m.GetMixerAttributes()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpClientConfigValidationError{
				field:  "MixerAttributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetForwardAttributes()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return HttpClientConfigValidationError{
				field:  "ForwardAttributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HttpClientConfigValidationError is the validation error returned by
// HttpClientConfig.Validate if the designated constraints aren't met.
type HttpClientConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpClientConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpClientConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpClientConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpClientConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpClientConfigValidationError) ErrorName() string { return "HttpClientConfigValidationError" }

// Error satisfies the builtin error interface
func (e HttpClientConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpClientConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpClientConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpClientConfigValidationError{}

// Validate checks the field values on TcpClientConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TcpClientConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTransport()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return TcpClientConfigValidationError{
				field:  "Transport",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMixerAttributes()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return TcpClientConfigValidationError{
				field:  "MixerAttributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DisableCheckCalls

	// no validation rules for DisableReportCalls

	if v, ok := interface{}(m.GetConnectionQuotaSpec()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return TcpClientConfigValidationError{
				field:  "ConnectionQuotaSpec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetReportInterval()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return TcpClientConfigValidationError{
				field:  "ReportInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TcpClientConfigValidationError is the validation error returned by
// TcpClientConfig.Validate if the designated constraints aren't met.
type TcpClientConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TcpClientConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TcpClientConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TcpClientConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TcpClientConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TcpClientConfigValidationError) ErrorName() string { return "TcpClientConfigValidationError" }

// Error satisfies the builtin error interface
func (e TcpClientConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpClientConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TcpClientConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TcpClientConfigValidationError{}
