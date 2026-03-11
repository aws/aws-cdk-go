//go:build no_runtime_type_checking

package previewawseventsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEventBusErrorLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnEventBusErrorLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnEventBusErrorLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnEventBusErrorLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnEventBusErrorLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnEventBusErrorLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnEventBusErrorLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnEventBusErrorLogsS3Props) error {
	return nil
}

