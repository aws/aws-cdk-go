//go:build no_runtime_type_checking

package previewawsapsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnScraperApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnScraperApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnScraperApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnScraperApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

