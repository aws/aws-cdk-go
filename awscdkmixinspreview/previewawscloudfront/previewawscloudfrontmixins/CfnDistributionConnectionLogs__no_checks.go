//go:build no_runtime_type_checking

package previewawscloudfrontmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDistributionConnectionLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnDistributionConnectionLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnDistributionConnectionLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnDistributionConnectionLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnDistributionConnectionLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnDistributionConnectionLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnDistributionConnectionLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnDistributionConnectionLogsS3Props) error {
	return nil
}

