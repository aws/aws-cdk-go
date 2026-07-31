//go:build no_runtime_type_checking

package awsmediaconnectalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateBridgeSourceConfiguration_FlowParameters(source *BridgeFlowSource) error {
	return nil
}

func validateBridgeSourceConfiguration_NetworkParameters(source *BridgeNetworkSource) error {
	return nil
}

func validateNewBridgeSourceConfigurationParameters(flowConfig *BridgeFlowSource, networkConfig *BridgeNetworkSource) error {
	return nil
}

