//go:build no_runtime_type_checking

package previewawsapigatewaymixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRestApiAccessLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnRestApiAccessLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnRestApiAccessLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnRestApiAccessLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnRestApiAccessLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnRestApiAccessLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnRestApiAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnRestApiAccessLogsS3Props) error {
	return nil
}

