//go:build no_runtime_type_checking

package previewawseksmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCapabilityEksCapabilityKroLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCapabilityEksCapabilityKroLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityKroLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCapabilityEksCapabilityKroLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityKroLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnCapabilityEksCapabilityKroLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityKroLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnCapabilityEksCapabilityKroLogsS3Props) error {
	return nil
}

