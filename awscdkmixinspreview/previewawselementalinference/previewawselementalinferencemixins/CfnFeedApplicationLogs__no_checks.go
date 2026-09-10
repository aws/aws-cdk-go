//go:build no_runtime_type_checking

package previewawselementalinferencemixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFeedApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnFeedApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnFeedApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnFeedApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnFeedApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnFeedApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnFeedApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnFeedApplicationLogsS3Props) error {
	return nil
}

