//go:build no_runtime_type_checking

package previewawsosismixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPipelinePipelineLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnPipelinePipelineLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnPipelinePipelineLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnPipelinePipelineLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnPipelinePipelineLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnPipelinePipelineLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnPipelinePipelineLogsS3Props) error {
	return nil
}

