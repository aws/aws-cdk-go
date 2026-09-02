//go:build no_runtime_type_checking

package previewawsbedrockmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAgentAliasApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnAgentAliasApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnAgentAliasApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnAgentAliasApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnAgentAliasApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnAgentAliasApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnAgentAliasApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnAgentAliasApplicationLogsS3Props) error {
	return nil
}

