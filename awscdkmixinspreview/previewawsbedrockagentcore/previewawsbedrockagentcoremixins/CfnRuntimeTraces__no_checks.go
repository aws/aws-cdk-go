//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRuntimeTraces) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnRuntimeTracesDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnRuntimeTraces) validateToXRayParameters(props *CfnRuntimeTracesXRayProps) error {
	return nil
}

