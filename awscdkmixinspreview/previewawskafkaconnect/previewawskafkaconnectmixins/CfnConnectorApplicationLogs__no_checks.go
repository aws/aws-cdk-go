//go:build no_runtime_type_checking

package previewawskafkaconnectmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnConnectorApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnConnectorApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnConnectorApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnConnectorApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnConnectorApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnConnectorApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnConnectorApplicationLogsS3Props) error {
	return nil
}

