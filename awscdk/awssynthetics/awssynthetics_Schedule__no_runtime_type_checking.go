//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awssynthetics

// Building without runtime type checking enabled, so all the below just return nil

func validateSchedule_CronParameters(options *CronOptions) error {
	return nil
}

func validateSchedule_ExpressionParameters(expression *string) error {
	return nil
}

func validateSchedule_RateParameters(interval awscdk.Duration) error {
	return nil
}

