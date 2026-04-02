//go:build !no_runtime_type_checking

package awsmediapackagev2alpha

import (
	"fmt"
)

func validateNumericExpression_RangeParameters(start *float64, end *float64) error {
	if start == nil {
		return fmt.Errorf("parameter start is required, but nil was provided")
	}

	if end == nil {
		return fmt.Errorf("parameter end is required, but nil was provided")
	}

	return nil
}

func validateNumericExpression_ValueParameters(v *float64) error {
	if v == nil {
		return fmt.Errorf("parameter v is required, but nil was provided")
	}

	return nil
}

func validateNewNumericExpressionParameters(_expression *string, _values *[]*float64) error {
	if _expression == nil {
		return fmt.Errorf("parameter _expression is required, but nil was provided")
	}

	if _values == nil {
		return fmt.Errorf("parameter _values is required, but nil was provided")
	}

	return nil
}

