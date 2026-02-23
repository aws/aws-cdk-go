//go:build no_runtime_type_checking

package previewawss3mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBucketS3ServerAccessLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnBucketS3ServerAccessLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnBucketS3ServerAccessLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnBucketS3ServerAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

