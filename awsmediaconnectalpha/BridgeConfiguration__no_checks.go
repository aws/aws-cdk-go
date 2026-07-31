//go:build no_runtime_type_checking

package awsmediaconnectalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateBridgeConfiguration_EgressParameters(config *EgressBridgeConfiguration) error {
	return nil
}

func validateBridgeConfiguration_IngressParameters(config *IngressBridgeConfiguration) error {
	return nil
}

