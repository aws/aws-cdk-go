//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func validateInstanceTypeFilter_AllowParameters(instanceTypes *[]awsec2.InstanceType) error {
	return nil
}

func validateInstanceTypeFilter_ExcludeParameters(instanceTypes *[]awsec2.InstanceType) error {
	return nil
}

