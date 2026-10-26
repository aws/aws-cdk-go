//go:build no_runtime_type_checking

package previewawsgameliftmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGameSessionQueueQueuePlacementLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnGameSessionQueueQueuePlacementLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnGameSessionQueueQueuePlacementLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnGameSessionQueueQueuePlacementLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnGameSessionQueueQueuePlacementLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnGameSessionQueueQueuePlacementLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnGameSessionQueueQueuePlacementLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnGameSessionQueueQueuePlacementLogsS3Props) error {
	return nil
}

