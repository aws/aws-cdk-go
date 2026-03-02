//go:build no_runtime_type_checking

package previewawsqbusinessmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnApplicationEventLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnApplicationEventLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationEventLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnApplicationEventLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationEventLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnApplicationEventLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationEventLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnApplicationEventLogsS3Props) error {
	return nil
}

