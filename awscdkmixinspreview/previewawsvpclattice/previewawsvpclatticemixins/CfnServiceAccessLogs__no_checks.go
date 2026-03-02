//go:build no_runtime_type_checking

package previewawsvpclatticemixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnServiceAccessLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnServiceAccessLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnServiceAccessLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnServiceAccessLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnServiceAccessLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnServiceAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnServiceAccessLogsS3Props) error {
	return nil
}

