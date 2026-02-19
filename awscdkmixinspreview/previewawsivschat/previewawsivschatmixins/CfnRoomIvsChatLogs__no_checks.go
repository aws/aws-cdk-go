//go:build no_runtime_type_checking

package previewawsivschatmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRoomIvsChatLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnRoomIvsChatLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnRoomIvsChatLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnRoomIvsChatLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

