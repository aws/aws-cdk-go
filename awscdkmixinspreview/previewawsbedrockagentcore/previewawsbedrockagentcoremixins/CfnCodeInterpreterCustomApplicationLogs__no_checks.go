//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCodeInterpreterCustomApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCodeInterpreterCustomApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterCustomApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCodeInterpreterCustomApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterCustomApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnCodeInterpreterCustomApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterCustomApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnCodeInterpreterCustomApplicationLogsS3Props) error {
	return nil
}

