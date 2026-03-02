//go:build no_runtime_type_checking

package previewawsmediapackagev2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnChannelGroupEgressAccessLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnChannelGroupEgressAccessLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnChannelGroupEgressAccessLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnChannelGroupEgressAccessLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnChannelGroupEgressAccessLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnChannelGroupEgressAccessLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnChannelGroupEgressAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnChannelGroupEgressAccessLogsS3Props) error {
	return nil
}

