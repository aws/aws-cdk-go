//go:build no_runtime_type_checking

package previewawsmediatailormixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPlaybackConfigurationManifestServiceLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnPlaybackConfigurationManifestServiceLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnPlaybackConfigurationManifestServiceLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnPlaybackConfigurationManifestServiceLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnPlaybackConfigurationManifestServiceLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnPlaybackConfigurationManifestServiceLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnPlaybackConfigurationManifestServiceLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnPlaybackConfigurationManifestServiceLogsS3Props) error {
	return nil
}

