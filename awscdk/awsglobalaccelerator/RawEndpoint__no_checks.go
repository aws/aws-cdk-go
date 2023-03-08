//go:build no_runtime_type_checking

package awsglobalaccelerator

// Building without runtime type checking enabled, so all the below just return nil

func validateNewRawEndpointParameters(props *RawEndpointProps) error {
	return nil
}

