//go:build no_runtime_type_checking

package previewawspcsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterPcsJobcompLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnClusterPcsJobcompLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterPcsJobcompLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnClusterPcsJobcompLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterPcsJobcompLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnClusterPcsJobcompLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterPcsJobcompLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnClusterPcsJobcompLogsS3Props) error {
	return nil
}

