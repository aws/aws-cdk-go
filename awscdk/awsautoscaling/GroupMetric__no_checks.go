//go:build no_runtime_type_checking

package awsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func validateNewGroupMetricParameters(name *string) error {
	return nil
}

