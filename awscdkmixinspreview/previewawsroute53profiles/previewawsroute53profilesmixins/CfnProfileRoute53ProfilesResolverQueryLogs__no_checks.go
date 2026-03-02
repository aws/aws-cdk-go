//go:build no_runtime_type_checking

package previewawsroute53profilesmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnProfileRoute53ProfilesResolverQueryLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnProfileRoute53ProfilesResolverQueryLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnProfileRoute53ProfilesResolverQueryLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnProfileRoute53ProfilesResolverQueryLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnProfileRoute53ProfilesResolverQueryLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnProfileRoute53ProfilesResolverQueryLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnProfileRoute53ProfilesResolverQueryLogsS3Props) error {
	return nil
}

