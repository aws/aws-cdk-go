//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRuntimeApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnRuntimeApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnRuntimeApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnRuntimeApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

