//go:build no_runtime_type_checking

package previewawseksmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterAutoModeComputeLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnClusterAutoModeComputeLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAutoModeComputeLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnClusterAutoModeComputeLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAutoModeComputeLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnClusterAutoModeComputeLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAutoModeComputeLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnClusterAutoModeComputeLogsS3Props) error {
	return nil
}

