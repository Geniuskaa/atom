// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: car.proto

package v1

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

// Validate checks the field values on CarInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CarInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CarInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CarInfoMultiError, or nil if none found.
func (m *CarInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *CarInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Plate

	// no validation rules for Model

	// no validation rules for YearOfManufacture

	// no validation rules for Type

	if len(errors) > 0 {
		return CarInfoMultiError(errors)
	}

	return nil
}

// CarInfoMultiError is an error wrapping multiple validation errors returned
// by CarInfo.ValidateAll() if the designated constraints aren't met.
type CarInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CarInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CarInfoMultiError) AllErrors() []error { return m }

// CarInfoValidationError is the validation error returned by CarInfo.Validate
// if the designated constraints aren't met.
type CarInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CarInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CarInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CarInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CarInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CarInfoValidationError) ErrorName() string { return "CarInfoValidationError" }

// Error satisfies the builtin error interface
func (e CarInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCarInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CarInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CarInfoValidationError{}

// Validate checks the field values on AllCarsInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AllCarsInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AllCarsInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AllCarsInfoMultiError, or
// nil if none found.
func (m *AllCarsInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *AllCarsInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCars() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AllCarsInfoValidationError{
						field:  fmt.Sprintf("Cars[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AllCarsInfoValidationError{
						field:  fmt.Sprintf("Cars[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AllCarsInfoValidationError{
					field:  fmt.Sprintf("Cars[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AllCarsInfoMultiError(errors)
	}

	return nil
}

// AllCarsInfoMultiError is an error wrapping multiple validation errors
// returned by AllCarsInfo.ValidateAll() if the designated constraints aren't met.
type AllCarsInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AllCarsInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AllCarsInfoMultiError) AllErrors() []error { return m }

// AllCarsInfoValidationError is the validation error returned by
// AllCarsInfo.Validate if the designated constraints aren't met.
type AllCarsInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AllCarsInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AllCarsInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AllCarsInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AllCarsInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AllCarsInfoValidationError) ErrorName() string { return "AllCarsInfoValidationError" }

// Error satisfies the builtin error interface
func (e AllCarsInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAllCarsInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AllCarsInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AllCarsInfoValidationError{}

// Validate checks the field values on UploadCarInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UploadCarInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadCarInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UploadCarInfoMultiError, or
// nil if none found.
func (m *UploadCarInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadCarInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Plate

	// no validation rules for Model

	// no validation rules for YearOfManufacture

	// no validation rules for Type

	if len(errors) > 0 {
		return UploadCarInfoMultiError(errors)
	}

	return nil
}

// UploadCarInfoMultiError is an error wrapping multiple validation errors
// returned by UploadCarInfo.ValidateAll() if the designated constraints
// aren't met.
type UploadCarInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadCarInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadCarInfoMultiError) AllErrors() []error { return m }

// UploadCarInfoValidationError is the validation error returned by
// UploadCarInfo.Validate if the designated constraints aren't met.
type UploadCarInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadCarInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadCarInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadCarInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadCarInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadCarInfoValidationError) ErrorName() string { return "UploadCarInfoValidationError" }

// Error satisfies the builtin error interface
func (e UploadCarInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadCarInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadCarInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadCarInfoValidationError{}

// Validate checks the field values on CarRegistered with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CarRegistered) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CarRegistered with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CarRegisteredMultiError, or
// nil if none found.
func (m *CarRegistered) ValidateAll() error {
	return m.validate(true)
}

func (m *CarRegistered) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for HasError

	if len(errors) > 0 {
		return CarRegisteredMultiError(errors)
	}

	return nil
}

// CarRegisteredMultiError is an error wrapping multiple validation errors
// returned by CarRegistered.ValidateAll() if the designated constraints
// aren't met.
type CarRegisteredMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CarRegisteredMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CarRegisteredMultiError) AllErrors() []error { return m }

// CarRegisteredValidationError is the validation error returned by
// CarRegistered.Validate if the designated constraints aren't met.
type CarRegisteredValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CarRegisteredValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CarRegisteredValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CarRegisteredValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CarRegisteredValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CarRegisteredValidationError) ErrorName() string { return "CarRegisteredValidationError" }

// Error satisfies the builtin error interface
func (e CarRegisteredValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCarRegistered.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CarRegisteredValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CarRegisteredValidationError{}

// Validate checks the field values on IsOK with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *IsOK) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IsOK with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in IsOKMultiError, or nil if none found.
func (m *IsOK) ValidateAll() error {
	return m.validate(true)
}

func (m *IsOK) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ok

	if len(errors) > 0 {
		return IsOKMultiError(errors)
	}

	return nil
}

// IsOKMultiError is an error wrapping multiple validation errors returned by
// IsOK.ValidateAll() if the designated constraints aren't met.
type IsOKMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IsOKMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IsOKMultiError) AllErrors() []error { return m }

// IsOKValidationError is the validation error returned by IsOK.Validate if the
// designated constraints aren't met.
type IsOKValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IsOKValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IsOKValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IsOKValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IsOKValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IsOKValidationError) ErrorName() string { return "IsOKValidationError" }

// Error satisfies the builtin error interface
func (e IsOKValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIsOK.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IsOKValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IsOKValidationError{}

// Validate checks the field values on CarID with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CarID) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CarID with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CarIDMultiError, or nil if none found.
func (m *CarID) ValidateAll() error {
	return m.validate(true)
}

func (m *CarID) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CarIDMultiError(errors)
	}

	return nil
}

// CarIDMultiError is an error wrapping multiple validation errors returned by
// CarID.ValidateAll() if the designated constraints aren't met.
type CarIDMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CarIDMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CarIDMultiError) AllErrors() []error { return m }

// CarIDValidationError is the validation error returned by CarID.Validate if
// the designated constraints aren't met.
type CarIDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CarIDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CarIDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CarIDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CarIDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CarIDValidationError) ErrorName() string { return "CarIDValidationError" }

// Error satisfies the builtin error interface
func (e CarIDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCarID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CarIDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CarIDValidationError{}

// Validate checks the field values on CarPlate with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CarPlate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CarPlate with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CarPlateMultiError, or nil
// if none found.
func (m *CarPlate) ValidateAll() error {
	return m.validate(true)
}

func (m *CarPlate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetPlate()); l < 7 || l > 9 {
		err := CarPlateValidationError{
			field:  "Plate",
			reason: "value length must be between 7 and 9 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CarPlateMultiError(errors)
	}

	return nil
}

// CarPlateMultiError is an error wrapping multiple validation errors returned
// by CarPlate.ValidateAll() if the designated constraints aren't met.
type CarPlateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CarPlateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CarPlateMultiError) AllErrors() []error { return m }

// CarPlateValidationError is the validation error returned by
// CarPlate.Validate if the designated constraints aren't met.
type CarPlateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CarPlateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CarPlateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CarPlateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CarPlateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CarPlateValidationError) ErrorName() string { return "CarPlateValidationError" }

// Error satisfies the builtin error interface
func (e CarPlateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCarPlate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CarPlateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CarPlateValidationError{}

// Validate checks the field values on CarPlateAndCosts with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CarPlateAndCosts) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CarPlateAndCosts with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CarPlateAndCostsMultiError, or nil if none found.
func (m *CarPlateAndCosts) ValidateAll() error {
	return m.validate(true)
}

func (m *CarPlateAndCosts) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Plate

	// no validation rules for FuelCosts

	// no validation rules for CarTax

	// no validation rules for AutoMaintenanceCosts

	// no validation rules for AutoConsumablesCost

	// no validation rules for InsuranceCost

	if len(errors) > 0 {
		return CarPlateAndCostsMultiError(errors)
	}

	return nil
}

// CarPlateAndCostsMultiError is an error wrapping multiple validation errors
// returned by CarPlateAndCosts.ValidateAll() if the designated constraints
// aren't met.
type CarPlateAndCostsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CarPlateAndCostsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CarPlateAndCostsMultiError) AllErrors() []error { return m }

// CarPlateAndCostsValidationError is the validation error returned by
// CarPlateAndCosts.Validate if the designated constraints aren't met.
type CarPlateAndCostsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CarPlateAndCostsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CarPlateAndCostsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CarPlateAndCostsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CarPlateAndCostsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CarPlateAndCostsValidationError) ErrorName() string { return "CarPlateAndCostsValidationError" }

// Error satisfies the builtin error interface
func (e CarPlateAndCostsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCarPlateAndCosts.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CarPlateAndCostsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CarPlateAndCostsValidationError{}

// Validate checks the field values on CarMileage with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CarMileage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CarMileage with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CarMileageMultiError, or
// nil if none found.
func (m *CarMileage) ValidateAll() error {
	return m.validate(true)
}

func (m *CarMileage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Mileage

	if len(errors) > 0 {
		return CarMileageMultiError(errors)
	}

	return nil
}

// CarMileageMultiError is an error wrapping multiple validation errors
// returned by CarMileage.ValidateAll() if the designated constraints aren't met.
type CarMileageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CarMileageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CarMileageMultiError) AllErrors() []error { return m }

// CarMileageValidationError is the validation error returned by
// CarMileage.Validate if the designated constraints aren't met.
type CarMileageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CarMileageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CarMileageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CarMileageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CarMileageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CarMileageValidationError) ErrorName() string { return "CarMileageValidationError" }

// Error satisfies the builtin error interface
func (e CarMileageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCarMileage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CarMileageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CarMileageValidationError{}

// Validate checks the field values on CarMileageUpload with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CarMileageUpload) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CarMileageUpload with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CarMileageUploadMultiError, or nil if none found.
func (m *CarMileageUpload) ValidateAll() error {
	return m.validate(true)
}

func (m *CarMileageUpload) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Mileage

	if len(errors) > 0 {
		return CarMileageUploadMultiError(errors)
	}

	return nil
}

// CarMileageUploadMultiError is an error wrapping multiple validation errors
// returned by CarMileageUpload.ValidateAll() if the designated constraints
// aren't met.
type CarMileageUploadMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CarMileageUploadMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CarMileageUploadMultiError) AllErrors() []error { return m }

// CarMileageUploadValidationError is the validation error returned by
// CarMileageUpload.Validate if the designated constraints aren't met.
type CarMileageUploadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CarMileageUploadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CarMileageUploadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CarMileageUploadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CarMileageUploadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CarMileageUploadValidationError) ErrorName() string { return "CarMileageUploadValidationError" }

// Error satisfies the builtin error interface
func (e CarMileageUploadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCarMileageUpload.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CarMileageUploadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CarMileageUploadValidationError{}

// Validate checks the field values on CarUsageCost with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CarUsageCost) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CarUsageCost with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CarUsageCostMultiError, or
// nil if none found.
func (m *CarUsageCost) ValidateAll() error {
	return m.validate(true)
}

func (m *CarUsageCost) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CostInRUB

	if len(errors) > 0 {
		return CarUsageCostMultiError(errors)
	}

	return nil
}

// CarUsageCostMultiError is an error wrapping multiple validation errors
// returned by CarUsageCost.ValidateAll() if the designated constraints aren't met.
type CarUsageCostMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CarUsageCostMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CarUsageCostMultiError) AllErrors() []error { return m }

// CarUsageCostValidationError is the validation error returned by
// CarUsageCost.Validate if the designated constraints aren't met.
type CarUsageCostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CarUsageCostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CarUsageCostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CarUsageCostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CarUsageCostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CarUsageCostValidationError) ErrorName() string { return "CarUsageCostValidationError" }

// Error satisfies the builtin error interface
func (e CarUsageCostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCarUsageCost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CarUsageCostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CarUsageCostValidationError{}

// Validate checks the field values on CarMetrics with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CarMetrics) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CarMetrics with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CarMetricsMultiError, or
// nil if none found.
func (m *CarMetrics) ValidateAll() error {
	return m.validate(true)
}

func (m *CarMetrics) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ID

	// no validation rules for Data

	if m.Period != nil {
		// no validation rules for Period
	}

	if len(errors) > 0 {
		return CarMetricsMultiError(errors)
	}

	return nil
}

// CarMetricsMultiError is an error wrapping multiple validation errors
// returned by CarMetrics.ValidateAll() if the designated constraints aren't met.
type CarMetricsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CarMetricsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CarMetricsMultiError) AllErrors() []error { return m }

// CarMetricsValidationError is the validation error returned by
// CarMetrics.Validate if the designated constraints aren't met.
type CarMetricsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CarMetricsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CarMetricsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CarMetricsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CarMetricsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CarMetricsValidationError) ErrorName() string { return "CarMetricsValidationError" }

// Error satisfies the builtin error interface
func (e CarMetricsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCarMetrics.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CarMetricsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CarMetricsValidationError{}
