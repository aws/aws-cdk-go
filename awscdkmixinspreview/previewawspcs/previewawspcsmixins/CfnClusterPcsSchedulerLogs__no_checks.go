//go:build no_runtime_type_checking

package previewawspcsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterPcsSchedulerLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnClusterPcsSchedulerLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterPcsSchedulerLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnClusterPcsSchedulerLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterPcsSchedulerLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnClusterPcsSchedulerLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterPcsSchedulerLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnClusterPcsSchedulerLogsS3Props) error {
	return nil
}

