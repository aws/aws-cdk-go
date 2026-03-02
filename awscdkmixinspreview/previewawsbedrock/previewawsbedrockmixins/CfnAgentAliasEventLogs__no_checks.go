//go:build no_runtime_type_checking

package previewawsbedrockmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAgentAliasEventLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnAgentAliasEventLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnAgentAliasEventLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnAgentAliasEventLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnAgentAliasEventLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnAgentAliasEventLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnAgentAliasEventLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnAgentAliasEventLogsS3Props) error {
	return nil
}

