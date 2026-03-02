//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMemoryTraces) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnMemoryTracesDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnMemoryTraces) validateToXRayParameters(props *CfnMemoryTracesXRayProps) error {
	return nil
}

