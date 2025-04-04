//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func validateTableThreshold_AboveParameters(value *float64) error {
	return nil
}

func validateTableThreshold_BelowParameters(value *float64) error {
	return nil
}

func validateTableThreshold_BetweenParameters(lowerBound *float64, upperBound *float64) error {
	return nil
}

