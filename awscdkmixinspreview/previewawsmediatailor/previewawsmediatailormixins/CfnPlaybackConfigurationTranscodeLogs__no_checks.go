//go:build no_runtime_type_checking

package previewawsmediatailormixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPlaybackConfigurationTranscodeLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnPlaybackConfigurationTranscodeLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnPlaybackConfigurationTranscodeLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnPlaybackConfigurationTranscodeLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnPlaybackConfigurationTranscodeLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnPlaybackConfigurationTranscodeLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnPlaybackConfigurationTranscodeLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnPlaybackConfigurationTranscodeLogsS3Props) error {
	return nil
}

