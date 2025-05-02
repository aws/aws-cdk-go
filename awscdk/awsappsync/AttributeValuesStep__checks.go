//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"
)

func (a *jsiiProxy_AttributeValuesStep) validateIsParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAttributeValuesStepParameters(attr *string, container *string, assignments *[]Assign) error {
	if attr == nil {
		return fmt.Errorf("parameter attr is required, but nil was provided")
	}

	if container == nil {
		return fmt.Errorf("parameter container is required, but nil was provided")
	}

	if assignments == nil {
		return fmt.Errorf("parameter assignments is required, but nil was provided")
	}

	return nil
}

