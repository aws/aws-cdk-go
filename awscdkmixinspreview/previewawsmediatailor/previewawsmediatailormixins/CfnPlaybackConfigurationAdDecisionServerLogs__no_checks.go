//go:build no_runtime_type_checking

package previewawsmediatailormixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPlaybackConfigurationAdDecisionServerLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnPlaybackConfigurationAdDecisionServerLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnPlaybackConfigurationAdDecisionServerLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnPlaybackConfigurationAdDecisionServerLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnPlaybackConfigurationAdDecisionServerLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnPlaybackConfigurationAdDecisionServerLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnPlaybackConfigurationAdDecisionServerLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnPlaybackConfigurationAdDecisionServerLogsS3Props) error {
	return nil
}

