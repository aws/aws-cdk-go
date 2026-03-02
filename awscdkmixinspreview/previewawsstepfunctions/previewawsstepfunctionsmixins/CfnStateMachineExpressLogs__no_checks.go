//go:build no_runtime_type_checking

package previewawsstepfunctionsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStateMachineExpressLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnStateMachineExpressLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnStateMachineExpressLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnStateMachineExpressLogsLogGroupProps) error {
	return nil
}

