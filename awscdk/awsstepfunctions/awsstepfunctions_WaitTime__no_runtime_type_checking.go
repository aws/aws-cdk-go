//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func validateWaitTime_DurationParameters(duration awscdk.Duration) error {
	return nil
}

func validateWaitTime_SecondsPathParameters(path *string) error {
	return nil
}

func validateWaitTime_TimestampParameters(timestamp *string) error {
	return nil
}

func validateWaitTime_TimestampPathParameters(path *string) error {
	return nil
}

