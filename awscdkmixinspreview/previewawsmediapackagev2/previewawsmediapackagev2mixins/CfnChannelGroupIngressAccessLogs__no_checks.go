//go:build no_runtime_type_checking

package previewawsmediapackagev2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnChannelGroupIngressAccessLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnChannelGroupIngressAccessLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnChannelGroupIngressAccessLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnChannelGroupIngressAccessLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnChannelGroupIngressAccessLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnChannelGroupIngressAccessLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnChannelGroupIngressAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnChannelGroupIngressAccessLogsS3Props) error {
	return nil
}

