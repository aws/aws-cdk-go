//go:build no_runtime_type_checking

package previewawsbedrockmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFlowApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnFlowApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnFlowApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnFlowApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnFlowApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnFlowApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnFlowApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnFlowApplicationLogsS3Props) error {
	return nil
}

