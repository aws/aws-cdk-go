//go:build no_runtime_type_checking

package previewawsb2bimixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTransformerB2biExecutionLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnTransformerB2biExecutionLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnTransformerB2biExecutionLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnTransformerB2biExecutionLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

