//go:build no_runtime_type_checking

package previewawssesmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMailManagerRuleSetApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnMailManagerRuleSetApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnMailManagerRuleSetApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnMailManagerRuleSetApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnMailManagerRuleSetApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnMailManagerRuleSetApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnMailManagerRuleSetApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnMailManagerRuleSetApplicationLogsS3Props) error {
	return nil
}

