//go:build no_runtime_type_checking

package previewawswisdommixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAssistantEventLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnAssistantEventLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnAssistantEventLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnAssistantEventLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

