//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCodeInterpreterCustomApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterCustomApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterCustomApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnCodeInterpreterCustomApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

