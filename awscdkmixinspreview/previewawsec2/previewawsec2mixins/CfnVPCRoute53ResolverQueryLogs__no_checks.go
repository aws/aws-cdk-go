//go:build no_runtime_type_checking

package previewawsec2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVPCRoute53ResolverQueryLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnVPCRoute53ResolverQueryLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnVPCRoute53ResolverQueryLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnVPCRoute53ResolverQueryLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnVPCRoute53ResolverQueryLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnVPCRoute53ResolverQueryLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnVPCRoute53ResolverQueryLogsS3Props) error {
	return nil
}

