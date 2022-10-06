//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsappsync

import (
	"fmt"
)

func (a *jsiiProxy_AttributeValues) validateAttributeParameters(attr *string) error {
	if attr == nil {
		return fmt.Errorf("parameter attr is required, but nil was provided")
	}

	return nil
}

func validateNewAttributeValuesParameters(container *string) error {
	if container == nil {
		return fmt.Errorf("parameter container is required, but nil was provided")
	}

	return nil
}

