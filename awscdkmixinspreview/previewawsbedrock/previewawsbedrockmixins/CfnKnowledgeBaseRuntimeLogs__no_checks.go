//go:build no_runtime_type_checking

package previewawsbedrockmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnKnowledgeBaseRuntimeLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnKnowledgeBaseRuntimeLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnKnowledgeBaseRuntimeLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnKnowledgeBaseRuntimeLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnKnowledgeBaseRuntimeLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnKnowledgeBaseRuntimeLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnKnowledgeBaseRuntimeLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnKnowledgeBaseRuntimeLogsS3Props) error {
	return nil
}

