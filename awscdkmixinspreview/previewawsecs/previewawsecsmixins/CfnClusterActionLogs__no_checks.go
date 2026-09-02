//go:build no_runtime_type_checking

package previewawsecsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterActionLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnClusterActionLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterActionLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnClusterActionLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterActionLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnClusterActionLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterActionLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnClusterActionLogsS3Props) error {
	return nil
}

