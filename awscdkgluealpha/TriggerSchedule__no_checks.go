//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateTriggerSchedule_CronParameters(options *awsevents.CronOptions) error {
	return nil
}

func validateTriggerSchedule_ExpressionParameters(expression *string) error {
	return nil
}

