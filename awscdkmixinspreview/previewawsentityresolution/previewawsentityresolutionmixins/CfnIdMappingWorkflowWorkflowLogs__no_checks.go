//go:build no_runtime_type_checking

package previewawsentityresolutionmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnIdMappingWorkflowWorkflowLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnIdMappingWorkflowWorkflowLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnIdMappingWorkflowWorkflowLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnIdMappingWorkflowWorkflowLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnIdMappingWorkflowWorkflowLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnIdMappingWorkflowWorkflowLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnIdMappingWorkflowWorkflowLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnIdMappingWorkflowWorkflowLogsS3Props) error {
	return nil
}

