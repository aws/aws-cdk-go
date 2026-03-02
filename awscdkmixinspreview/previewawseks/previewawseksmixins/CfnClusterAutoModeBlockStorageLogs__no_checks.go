//go:build no_runtime_type_checking

package previewawseksmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterAutoModeBlockStorageLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnClusterAutoModeBlockStorageLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAutoModeBlockStorageLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnClusterAutoModeBlockStorageLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAutoModeBlockStorageLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnClusterAutoModeBlockStorageLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAutoModeBlockStorageLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnClusterAutoModeBlockStorageLogsS3Props) error {
	return nil
}

