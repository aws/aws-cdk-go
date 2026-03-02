//go:build no_runtime_type_checking

package previewawss3mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBucketS3ServerAccessLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnBucketS3ServerAccessLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnBucketS3ServerAccessLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnBucketS3ServerAccessLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnBucketS3ServerAccessLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnBucketS3ServerAccessLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnBucketS3ServerAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnBucketS3ServerAccessLogsS3Props) error {
	return nil
}

