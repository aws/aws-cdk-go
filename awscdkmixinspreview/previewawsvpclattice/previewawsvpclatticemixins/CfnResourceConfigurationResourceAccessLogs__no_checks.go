//go:build no_runtime_type_checking

package previewawsvpclatticemixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnResourceConfigurationResourceAccessLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnResourceConfigurationResourceAccessLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnResourceConfigurationResourceAccessLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnResourceConfigurationResourceAccessLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnResourceConfigurationResourceAccessLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnResourceConfigurationResourceAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnResourceConfigurationResourceAccessLogsS3Props) error {
	return nil
}

