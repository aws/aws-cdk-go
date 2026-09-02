//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWorkloadIdentityTraces) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnWorkloadIdentityTracesDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnWorkloadIdentityTraces) validateToXRayParameters(props *CfnWorkloadIdentityTracesXRayProps) error {
	return nil
}

