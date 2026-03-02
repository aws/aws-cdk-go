//go:build no_runtime_type_checking

package previewawsb2bimixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTransformerB2biExecutionLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnTransformerB2biExecutionLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnTransformerB2biExecutionLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnTransformerB2biExecutionLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnTransformerB2biExecutionLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnTransformerB2biExecutionLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnTransformerB2biExecutionLogsS3Props) error {
	return nil
}

