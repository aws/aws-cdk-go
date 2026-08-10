//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCodeInterpreterUsageLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCodeInterpreterUsageLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterUsageLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCodeInterpreterUsageLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterUsageLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnCodeInterpreterUsageLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterUsageLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnCodeInterpreterUsageLogsS3Props) error {
	return nil
}

