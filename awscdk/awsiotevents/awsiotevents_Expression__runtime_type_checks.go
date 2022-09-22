//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsiotevents

import (
	"fmt"
)

func validateExpression_AndParameters(left Expression, right Expression) error {
	if left == nil {
		return fmt.Errorf("parameter left is required, but nil was provided")
	}

	if right == nil {
		return fmt.Errorf("parameter right is required, but nil was provided")
	}

	return nil
}

func validateExpression_CurrentInputParameters(input IInput) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}

	return nil
}

func validateExpression_EqParameters(left Expression, right Expression) error {
	if left == nil {
		return fmt.Errorf("parameter left is required, but nil was provided")
	}

	if right == nil {
		return fmt.Errorf("parameter right is required, but nil was provided")
	}

	return nil
}

func validateExpression_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateExpression_GtParameters(left Expression, right Expression) error {
	if left == nil {
		return fmt.Errorf("parameter left is required, but nil was provided")
	}

	if right == nil {
		return fmt.Errorf("parameter right is required, but nil was provided")
	}

	return nil
}

func validateExpression_GteParameters(left Expression, right Expression) error {
	if left == nil {
		return fmt.Errorf("parameter left is required, but nil was provided")
	}

	if right == nil {
		return fmt.Errorf("parameter right is required, but nil was provided")
	}

	return nil
}

func validateExpression_InputAttributeParameters(input IInput, path *string) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}

	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateExpression_LtParameters(left Expression, right Expression) error {
	if left == nil {
		return fmt.Errorf("parameter left is required, but nil was provided")
	}

	if right == nil {
		return fmt.Errorf("parameter right is required, but nil was provided")
	}

	return nil
}

func validateExpression_LteParameters(left Expression, right Expression) error {
	if left == nil {
		return fmt.Errorf("parameter left is required, but nil was provided")
	}

	if right == nil {
		return fmt.Errorf("parameter right is required, but nil was provided")
	}

	return nil
}

func validateExpression_NeqParameters(left Expression, right Expression) error {
	if left == nil {
		return fmt.Errorf("parameter left is required, but nil was provided")
	}

	if right == nil {
		return fmt.Errorf("parameter right is required, but nil was provided")
	}

	return nil
}

func validateExpression_OrParameters(left Expression, right Expression) error {
	if left == nil {
		return fmt.Errorf("parameter left is required, but nil was provided")
	}

	if right == nil {
		return fmt.Errorf("parameter right is required, but nil was provided")
	}

	return nil
}

