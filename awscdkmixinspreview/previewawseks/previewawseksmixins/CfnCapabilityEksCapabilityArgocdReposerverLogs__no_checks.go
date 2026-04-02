//go:build no_runtime_type_checking

package previewawseksmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdReposerverLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCapabilityEksCapabilityArgocdReposerverLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdReposerverLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCapabilityEksCapabilityArgocdReposerverLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdReposerverLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnCapabilityEksCapabilityArgocdReposerverLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityArgocdReposerverLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnCapabilityEksCapabilityArgocdReposerverLogsS3Props) error {
	return nil
}

