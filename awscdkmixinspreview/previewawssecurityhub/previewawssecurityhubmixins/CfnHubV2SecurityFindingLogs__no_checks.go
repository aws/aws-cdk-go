//go:build no_runtime_type_checking

package previewawssecurityhubmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHubV2SecurityFindingLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnHubV2SecurityFindingLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnHubV2SecurityFindingLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnHubV2SecurityFindingLogsLogGroupProps) error {
	return nil
}

