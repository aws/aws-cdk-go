//go:build no_runtime_type_checking

package previewawsconnectmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInstanceAmazonConnectFlowLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnInstanceAmazonConnectFlowLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnInstanceAmazonConnectFlowLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnInstanceAmazonConnectFlowLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnInstanceAmazonConnectFlowLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnInstanceAmazonConnectFlowLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnInstanceAmazonConnectFlowLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnInstanceAmazonConnectFlowLogsS3Props) error {
	return nil
}

