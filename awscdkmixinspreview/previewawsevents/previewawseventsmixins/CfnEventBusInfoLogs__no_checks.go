//go:build no_runtime_type_checking

package previewawseventsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEventBusInfoLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnEventBusInfoLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnEventBusInfoLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnEventBusInfoLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnEventBusInfoLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnEventBusInfoLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnEventBusInfoLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnEventBusInfoLogsS3Props) error {
	return nil
}

