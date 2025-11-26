//go:build no_runtime_type_checking

package previewawsmediapackagev2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnChannelGroupEgressAccessLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnChannelGroupEgressAccessLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnChannelGroupEgressAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

