//go:build no_runtime_type_checking

package previewawsapigatewaymixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRestApiExecutionLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnRestApiExecutionLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnRestApiExecutionLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnRestApiExecutionLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnRestApiExecutionLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnRestApiExecutionLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnRestApiExecutionLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnRestApiExecutionLogsS3Props) error {
	return nil
}

