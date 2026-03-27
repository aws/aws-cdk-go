//go:build no_runtime_type_checking

package previewawsroute53globalresolvermixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGlobalResolverGlobalResolverLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnGlobalResolverGlobalResolverLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnGlobalResolverGlobalResolverLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnGlobalResolverGlobalResolverLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnGlobalResolverGlobalResolverLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnGlobalResolverGlobalResolverLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnGlobalResolverGlobalResolverLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnGlobalResolverGlobalResolverLogsS3Props) error {
	return nil
}

