//go:build no_runtime_type_checking

package previewawsshieldmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnProtectionFlowLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnProtectionFlowLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnProtectionFlowLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnProtectionFlowLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnProtectionFlowLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnProtectionFlowLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnProtectionFlowLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnProtectionFlowLogsS3Props) error {
	return nil
}

