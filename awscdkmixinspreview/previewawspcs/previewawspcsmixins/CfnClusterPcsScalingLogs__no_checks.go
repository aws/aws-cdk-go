//go:build no_runtime_type_checking

package previewawspcsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterPcsScalingLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnClusterPcsScalingLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterPcsScalingLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnClusterPcsScalingLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterPcsScalingLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnClusterPcsScalingLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterPcsScalingLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnClusterPcsScalingLogsS3Props) error {
	return nil
}

