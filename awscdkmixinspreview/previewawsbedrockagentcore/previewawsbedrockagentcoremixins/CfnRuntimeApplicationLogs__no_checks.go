//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRuntimeApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnRuntimeApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnRuntimeApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnRuntimeApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnRuntimeApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnRuntimeApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnRuntimeApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnRuntimeApplicationLogsS3Props) error {
	return nil
}

