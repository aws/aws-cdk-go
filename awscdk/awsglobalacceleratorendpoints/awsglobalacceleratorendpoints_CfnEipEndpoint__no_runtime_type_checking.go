//go:build no_runtime_type_checking

package awsglobalacceleratorendpoints

// Building without runtime type checking enabled, so all the below just return nil

func validateNewCfnEipEndpointParameters(eip awsec2.CfnEIP, options *CfnEipEndpointProps) error {
	return nil
}

