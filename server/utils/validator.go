package utils

import (
	"errors"
	"reflect"
)

func ValidatePayload(body map[string]interface{}) error {
	// Requiredfields and it's types
	requiredFields := map[string]reflect.Kind{
		"title":  reflect.String,
		"author": reflect.String,
		"year":   reflect.Float64,
	}

	// Validate required fields
	err := validateRequiredFields(requiredFields, body)
	if err != nil {
		return err
	}

	// Validate types
	err = validateTypes(requiredFields, body)
	if err != nil {
		return err
	}

	// Validate non-empty fields
	err = validateContent(body)
	if err != nil {
		return err
	}

	return nil
}

func validateRequiredFields(requiredFields map[string]reflect.Kind, body map[string]interface{}) error {
	for k := range requiredFields {
		if contains(body, k) {
			continue
		} else {
			return errors.New("missing required field '" + k + "'")
		}
	}
	return nil
}

func contains(body map[string]interface{}, str string) bool {
	for k := range body {
		if k == str {
			return true
		}
	}
	return false
}

func validateTypes(bookRules map[string]reflect.Kind, body map[string]interface{}) error {
	for k, v := range body {
		if reflect.TypeOf(v).Kind() != bookRules[k] {
			return errors.New("field " + k + " must be of type " + bookRules[k].String())
		}
	}
	return nil
}

func validateContent(body map[string]interface{}) error {
	for k, v := range body {
		if v == "" {
			return errors.New("field " + k + " can not be empty")
		}
	}
	return nil
}
