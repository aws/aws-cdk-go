//go:build no_runtime_type_checking

package awsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func validateHealthChecks_Ec2Parameters(options *Ec2HealthChecksOptions) error {
	return nil
}

func validateHealthChecks_WithAdditionalChecksParameters(options *AdditionalHealthChecksOptions) error {
	return nil
}

