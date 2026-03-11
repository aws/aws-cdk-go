//go:build no_runtime_type_checking

package previewawselasticloadbalancingv2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLoadBalancerAlbHealthCheckLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnLoadBalancerAlbHealthCheckLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnLoadBalancerAlbHealthCheckLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnLoadBalancerAlbHealthCheckLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnLoadBalancerAlbHealthCheckLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnLoadBalancerAlbHealthCheckLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnLoadBalancerAlbHealthCheckLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnLoadBalancerAlbHealthCheckLogsS3Props) error {
	return nil
}

