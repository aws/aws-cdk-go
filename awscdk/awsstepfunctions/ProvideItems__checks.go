//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"
)

func validateProvideItems_JsonArrayParameters(array *[]interface{}) error {
	if array == nil {
		return fmt.Errorf("parameter array is required, but nil was provided")
	}

	return nil
}

func validateProvideItems_JsonataParameters(jsonataExpression *string) error {
	if jsonataExpression == nil {
		return fmt.Errorf("parameter jsonataExpression is required, but nil was provided")
	}

	return nil
}

