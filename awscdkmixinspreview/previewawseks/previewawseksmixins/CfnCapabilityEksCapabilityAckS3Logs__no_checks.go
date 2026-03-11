//go:build no_runtime_type_checking

package previewawseksmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCapabilityEksCapabilityAckS3Logs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnCapabilityEksCapabilityAckS3LogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityAckS3Logs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCapabilityEksCapabilityAckS3LogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityAckS3Logs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnCapabilityEksCapabilityAckS3LogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnCapabilityEksCapabilityAckS3Logs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnCapabilityEksCapabilityAckS3LogsS3Props) error {
	return nil
}

