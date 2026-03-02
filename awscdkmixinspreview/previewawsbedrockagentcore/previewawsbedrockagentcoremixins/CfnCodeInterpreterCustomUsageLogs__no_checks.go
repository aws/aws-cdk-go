//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCodeInterpreterCustomUsageLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCodeInterpreterCustomUsageLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterCustomUsageLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCodeInterpreterCustomUsageLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterCustomUsageLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnCodeInterpreterCustomUsageLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterCustomUsageLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnCodeInterpreterCustomUsageLogsS3Props) error {
	return nil
}

