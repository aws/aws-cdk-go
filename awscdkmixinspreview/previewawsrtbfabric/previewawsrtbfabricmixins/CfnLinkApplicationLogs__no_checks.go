//go:build no_runtime_type_checking

package previewawsrtbfabricmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLinkApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnLinkApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnLinkApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnLinkApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnLinkApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnLinkApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnLinkApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnLinkApplicationLogsS3Props) error {
	return nil
}

