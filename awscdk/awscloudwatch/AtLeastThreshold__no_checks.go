//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func validateAtLeastThreshold_CountParameters(count *float64) error {
	return nil
}

func validateAtLeastThreshold_PercentageParameters(percentage *float64) error {
	return nil
}

