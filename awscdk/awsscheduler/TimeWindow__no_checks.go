//go:build no_runtime_type_checking

package awsscheduler

// Building without runtime type checking enabled, so all the below just return nil

func validateTimeWindow_FlexibleParameters(maxWindow awscdk.Duration) error {
	return nil
}

