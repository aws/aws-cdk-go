//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func validateEc2MaxInstances_PercentageParameters(percentage *float64) error {
	return nil
}

func validateEc2MaxInstances_TargetsParameters(targets *float64) error {
	return nil
}

