//go:build no_runtime_type_checking

package previewawsentityresolutionmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMatchingWorkflowWorkflowLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnMatchingWorkflowWorkflowLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnMatchingWorkflowWorkflowLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnMatchingWorkflowWorkflowLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnMatchingWorkflowWorkflowLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnMatchingWorkflowWorkflowLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnMatchingWorkflowWorkflowLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnMatchingWorkflowWorkflowLogsS3Props) error {
	return nil
}

