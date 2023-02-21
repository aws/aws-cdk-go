//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func validateTimeout_AtParameters(path *string) error {
	return nil
}

func validateTimeout_DurationParameters(duration awscdk.Duration) error {
	return nil
}

