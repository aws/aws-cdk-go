//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func validateCapacity_AutoscaledParameters(options *AutoscaledCapacityOptions) error {
	return nil
}

func validateCapacity_FixedParameters(iops *float64) error {
	return nil
}

