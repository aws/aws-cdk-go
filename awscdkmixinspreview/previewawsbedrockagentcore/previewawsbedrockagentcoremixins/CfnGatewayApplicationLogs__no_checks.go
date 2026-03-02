//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGatewayApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnGatewayApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnGatewayApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnGatewayApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnGatewayApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnGatewayApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnGatewayApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnGatewayApplicationLogsS3Props) error {
	return nil
}

