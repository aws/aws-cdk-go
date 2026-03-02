//go:build no_runtime_type_checking

package previewawswisdommixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAssistantEventLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnAssistantEventLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnAssistantEventLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnAssistantEventLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnAssistantEventLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnAssistantEventLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnAssistantEventLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnAssistantEventLogsS3Props) error {
	return nil
}

