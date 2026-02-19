//go:build no_runtime_type_checking

package previewawscloudfrontmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDistributionConnectionLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnDistributionConnectionLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnDistributionConnectionLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnDistributionConnectionLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

