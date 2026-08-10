//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func validateScheduleExpression_AtParameters(date *CalendarDateTime) error {
	return nil
}

func validateScheduleExpression_CronParameters(options *CronOptions) error {
	return nil
}

func validateScheduleExpression_ExpressionParameters(expression *string) error {
	return nil
}

