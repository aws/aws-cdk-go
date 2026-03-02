//go:build no_runtime_type_checking

package previewawsbedrockmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAgentApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnAgentApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnAgentApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnAgentApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnAgentApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnAgentApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnAgentApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnAgentApplicationLogsS3Props) error {
	return nil
}

