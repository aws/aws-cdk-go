//go:build no_runtime_type_checking

package previewawsmediatailormixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPlaybackConfigurationTranscodeLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnPlaybackConfigurationTranscodeLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnPlaybackConfigurationTranscodeLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnPlaybackConfigurationTranscodeLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

