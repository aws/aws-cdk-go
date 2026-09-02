//go:build no_runtime_type_checking

package previewawssecurityhubmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHubSecurityFindingLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnHubSecurityFindingLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnHubSecurityFindingLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnHubSecurityFindingLogsLogGroupProps) error {
	return nil
}

