//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBrowserCustomUsageLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnBrowserCustomUsageLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnBrowserCustomUsageLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnBrowserCustomUsageLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnBrowserCustomUsageLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnBrowserCustomUsageLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnBrowserCustomUsageLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnBrowserCustomUsageLogsS3Props) error {
	return nil
}

