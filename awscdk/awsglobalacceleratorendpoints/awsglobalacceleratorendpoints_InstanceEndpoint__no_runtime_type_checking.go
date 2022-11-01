//go:build no_runtime_type_checking

package awsglobalacceleratorendpoints

// Building without runtime type checking enabled, so all the below just return nil

func validateNewInstanceEndpointParameters(instance awsec2.IInstance, options *InstanceEndpointProps) error {
	return nil
}

