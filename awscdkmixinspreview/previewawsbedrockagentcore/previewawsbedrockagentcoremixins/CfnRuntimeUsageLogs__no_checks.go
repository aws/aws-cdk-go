//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRuntimeUsageLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnRuntimeUsageLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnRuntimeUsageLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnRuntimeUsageLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnRuntimeUsageLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnRuntimeUsageLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnRuntimeUsageLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnRuntimeUsageLogsS3Props) error {
	return nil
}

