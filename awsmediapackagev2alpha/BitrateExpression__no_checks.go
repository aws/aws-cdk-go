//go:build no_runtime_type_checking

package awsmediapackagev2alpha

// Building without runtime type checking enabled, so all the below just return nil

func validateBitrateExpression_RangeParameters(start awscdk.Bitrate, end awscdk.Bitrate) error {
	return nil
}

func validateBitrateExpression_ValueParameters(v awscdk.Bitrate) error {
	return nil
}

func validateNewBitrateExpressionParameters(_expression *string, _values *[]*float64) error {
	return nil
}

