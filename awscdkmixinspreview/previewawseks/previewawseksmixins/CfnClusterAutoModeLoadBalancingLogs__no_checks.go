//go:build no_runtime_type_checking

package previewawseksmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterAutoModeLoadBalancingLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnClusterAutoModeLoadBalancingLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAutoModeLoadBalancingLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnClusterAutoModeLoadBalancingLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAutoModeLoadBalancingLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnClusterAutoModeLoadBalancingLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAutoModeLoadBalancingLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnClusterAutoModeLoadBalancingLogsS3Props) error {
	return nil
}

