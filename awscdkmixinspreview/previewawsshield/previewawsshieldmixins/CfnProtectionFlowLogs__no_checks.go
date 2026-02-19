//go:build no_runtime_type_checking

package previewawsshieldmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnProtectionFlowLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnProtectionFlowLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnProtectionFlowLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnProtectionFlowLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

