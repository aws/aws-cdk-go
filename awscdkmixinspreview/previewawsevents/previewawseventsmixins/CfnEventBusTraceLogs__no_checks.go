//go:build no_runtime_type_checking

package previewawseventsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEventBusTraceLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnEventBusTraceLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnEventBusTraceLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnEventBusTraceLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnEventBusTraceLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnEventBusTraceLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnEventBusTraceLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnEventBusTraceLogsS3Props) error {
	return nil
}

