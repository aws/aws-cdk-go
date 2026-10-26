//go:build no_runtime_type_checking

package previewawseksmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdNotificationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCapabilityEksCapabilityArgocdNotificationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdNotificationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCapabilityEksCapabilityArgocdNotificationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdNotificationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnCapabilityEksCapabilityArgocdNotificationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdNotificationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnCapabilityEksCapabilityArgocdNotificationLogsS3Props) error {
	return nil
}

