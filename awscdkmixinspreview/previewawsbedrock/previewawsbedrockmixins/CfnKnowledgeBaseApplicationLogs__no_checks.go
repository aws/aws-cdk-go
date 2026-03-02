//go:build no_runtime_type_checking

package previewawsbedrockmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnKnowledgeBaseApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnKnowledgeBaseApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnKnowledgeBaseApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnKnowledgeBaseApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnKnowledgeBaseApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnKnowledgeBaseApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnKnowledgeBaseApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnKnowledgeBaseApplicationLogsS3Props) error {
	return nil
}

