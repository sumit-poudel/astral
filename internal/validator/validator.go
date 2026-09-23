package validator

import (
	"regexp"
	"strings"
)

type Validator struct {
	FieldErrors map[string]string
}

func (v *Validator) IsValid() bool {
	if len(v.FieldErrors) == 0 {
		return true
	}
	return false
}

var EmailRx = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

func (v *Validator) CheckFieldWithErr(f func() (bool, error), key, message string) error {
	if ok, err := f(); err != nil {
		return err
	} else if !ok {
		v.AddFieldError(key, message)
	}
	return nil
}

func (v *Validator) NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func (v *Validator) Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func (v *Validator) IsEquals(value1 string, value2 string) bool {
	if strings.Compare(value1, value2) == 0 {
		return true
	}
	return false
}
