//go:build no_runtime_type_checking

package previewawsstepfunctionsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStateMachineStandardLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnStateMachineStandardLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnStateMachineStandardLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnStateMachineStandardLogsLogGroupProps) error {
	return nil
}

