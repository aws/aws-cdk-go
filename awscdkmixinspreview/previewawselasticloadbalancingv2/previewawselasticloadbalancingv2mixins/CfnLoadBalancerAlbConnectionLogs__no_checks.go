//go:build no_runtime_type_checking

package previewawselasticloadbalancingv2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLoadBalancerAlbConnectionLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnLoadBalancerAlbConnectionLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnLoadBalancerAlbConnectionLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnLoadBalancerAlbConnectionLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnLoadBalancerAlbConnectionLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnLoadBalancerAlbConnectionLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnLoadBalancerAlbConnectionLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnLoadBalancerAlbConnectionLogsS3Props) error {
	return nil
}

