//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBrowserCustomTraces) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnBrowserCustomTracesDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnBrowserCustomTraces) validateToXRayParameters(props *CfnBrowserCustomTracesXRayProps) error {
	return nil
}

