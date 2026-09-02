//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCodeInterpreterCustomTraces) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCodeInterpreterCustomTracesDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterCustomTraces) validateToXRayParameters(props *CfnCodeInterpreterCustomTracesXRayProps) error {
	return nil
}

