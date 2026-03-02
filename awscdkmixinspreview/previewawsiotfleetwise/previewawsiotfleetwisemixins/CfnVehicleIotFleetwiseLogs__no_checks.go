//go:build no_runtime_type_checking

package previewawsiotfleetwisemixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVehicleIotFleetwiseLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnVehicleIotFleetwiseLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnVehicleIotFleetwiseLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnVehicleIotFleetwiseLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnVehicleIotFleetwiseLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnVehicleIotFleetwiseLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnVehicleIotFleetwiseLogsS3Props) error {
	return nil
}

