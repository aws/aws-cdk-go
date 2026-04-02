//go:build no_runtime_type_checking

package awsmediapackagev2alpha

// Building without runtime type checking enabled, so all the below just return nil

func validateNumericExpression_RangeParameters(start *float64, end *float64) error {
	return nil
}

func validateNumericExpression_ValueParameters(v *float64) error {
	return nil
}

func validateNewNumericExpressionParameters(_expression *string, _values *[]*float64) error {
	return nil
}

