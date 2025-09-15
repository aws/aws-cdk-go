//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func validateNewPortMapParameters(networkmode NetworkMode, pm *PortMapping) error {
	return nil
}

