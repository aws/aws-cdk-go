//go:build no_runtime_type_checking

package previewawseksmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterAutoModeIpamLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnClusterAutoModeIpamLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAutoModeIpamLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnClusterAutoModeIpamLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAutoModeIpamLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnClusterAutoModeIpamLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAutoModeIpamLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnClusterAutoModeIpamLogsS3Props) error {
	return nil
}

