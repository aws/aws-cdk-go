//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGatewayTraces) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnGatewayTracesDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnGatewayTraces) validateToXRayParameters(props *CfnGatewayTracesXRayProps) error {
	return nil
}

