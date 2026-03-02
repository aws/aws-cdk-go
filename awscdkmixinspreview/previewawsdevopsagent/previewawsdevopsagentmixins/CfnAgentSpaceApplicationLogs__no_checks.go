//go:build no_runtime_type_checking

package previewawsdevopsagentmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAgentSpaceApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnAgentSpaceApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnAgentSpaceApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnAgentSpaceApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnAgentSpaceApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnAgentSpaceApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnAgentSpaceApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnAgentSpaceApplicationLogsS3Props) error {
	return nil
}

