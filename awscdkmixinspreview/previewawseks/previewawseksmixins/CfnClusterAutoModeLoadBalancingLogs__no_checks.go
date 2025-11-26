//go:build no_runtime_type_checking

package previewawseksmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterAutoModeLoadBalancingLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAutoModeLoadBalancingLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAutoModeLoadBalancingLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

