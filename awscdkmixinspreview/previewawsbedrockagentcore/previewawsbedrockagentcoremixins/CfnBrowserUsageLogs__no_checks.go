//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBrowserUsageLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnBrowserUsageLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnBrowserUsageLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnBrowserUsageLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnBrowserUsageLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnBrowserUsageLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnBrowserUsageLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnBrowserUsageLogsS3Props) error {
	return nil
}

