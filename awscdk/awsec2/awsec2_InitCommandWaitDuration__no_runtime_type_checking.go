//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateInitCommandWaitDuration_OfParameters(duration awscdk.Duration) error {
	return nil
}

