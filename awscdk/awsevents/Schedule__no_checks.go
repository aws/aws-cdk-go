//go:build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func validateSchedule_CronParameters(options *CronOptions) error {
	return nil
}

func validateSchedule_ExpressionParameters(expression *string) error {
	return nil
}

func validateSchedule_RateParameters(duration awscdk.Duration) error {
	return nil
}

