//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMemoryApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnMemoryApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnMemoryApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnMemoryApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnMemoryApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnMemoryApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnMemoryApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnMemoryApplicationLogsS3Props) error {
	return nil
}

