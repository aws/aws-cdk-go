//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateSchedule_ProtectedAtParameters(date *time.Time) error {
	return nil
}

func validateSchedule_ProtectedCronParameters(options *CronOptions) error {
	return nil
}

func validateSchedule_ProtectedExpressionParameters(expression *string) error {
	return nil
}

func validateSchedule_ProtectedRateParameters(duration Duration) error {
	return nil
}

