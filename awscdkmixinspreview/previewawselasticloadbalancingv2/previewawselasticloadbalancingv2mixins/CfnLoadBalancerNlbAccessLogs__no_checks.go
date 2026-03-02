//go:build no_runtime_type_checking

package previewawselasticloadbalancingv2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLoadBalancerNlbAccessLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnLoadBalancerNlbAccessLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnLoadBalancerNlbAccessLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnLoadBalancerNlbAccessLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnLoadBalancerNlbAccessLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnLoadBalancerNlbAccessLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnLoadBalancerNlbAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnLoadBalancerNlbAccessLogsS3Props) error {
	return nil
}

