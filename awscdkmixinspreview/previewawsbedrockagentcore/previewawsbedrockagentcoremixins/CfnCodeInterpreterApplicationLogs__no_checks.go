//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCodeInterpreterApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCodeInterpreterApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCodeInterpreterApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnCodeInterpreterApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnCodeInterpreterApplicationLogsS3Props) error {
	return nil
}

