//go:build no_runtime_type_checking

package awsapplicationautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func validateSchedule_AtParameters(moment *time.Time) error {
	return nil
}

func validateSchedule_CronParameters(options *awscdk.CronOptions) error {
	return nil
}

func validateSchedule_ExpressionParameters(expression *string) error {
	return nil
}

func validateSchedule_ProtectedAtParameters(date *time.Time) error {
	return nil
}

func validateSchedule_ProtectedCronParameters(options *awscdk.CronOptions) error {
	return nil
}

func validateSchedule_ProtectedExpressionParameters(expression *string) error {
	return nil
}

func validateSchedule_ProtectedRateParameters(duration awscdk.Duration) error {
	return nil
}

func validateSchedule_RateParameters(duration awscdk.Duration) error {
	return nil
}

