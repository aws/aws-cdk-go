//go:build no_runtime_type_checking

package awscdkscheduleralpha

// Building without runtime type checking enabled, so all the below just return nil

func validateScheduleExpression_AtParameters(date *time.Time) error {
	return nil
}

func validateScheduleExpression_CronParameters(options *awscdk.CronOptions) error {
	return nil
}

func validateScheduleExpression_ExpressionParameters(expression *string) error {
	return nil
}

func validateScheduleExpression_ProtectedAtParameters(date *time.Time) error {
	return nil
}

func validateScheduleExpression_ProtectedCronParameters(options *awscdk.CronOptions) error {
	return nil
}

func validateScheduleExpression_ProtectedExpressionParameters(expression *string) error {
	return nil
}

func validateScheduleExpression_ProtectedRateParameters(duration awscdk.Duration) error {
	return nil
}

func validateScheduleExpression_RateParameters(duration awscdk.Duration) error {
	return nil
}

