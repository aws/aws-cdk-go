//go:build no_runtime_type_checking

package previewawseksmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnCapabilityEksCapabilityArgocdApplicationLogsS3Props) error {
	return nil
}

