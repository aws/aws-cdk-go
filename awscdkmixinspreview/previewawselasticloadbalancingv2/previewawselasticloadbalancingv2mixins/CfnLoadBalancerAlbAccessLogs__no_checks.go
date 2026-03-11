//go:build no_runtime_type_checking

package previewawselasticloadbalancingv2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLoadBalancerAlbAccessLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnLoadBalancerAlbAccessLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnLoadBalancerAlbAccessLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnLoadBalancerAlbAccessLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnLoadBalancerAlbAccessLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnLoadBalancerAlbAccessLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnLoadBalancerAlbAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnLoadBalancerAlbAccessLogsS3Props) error {
	return nil
}

