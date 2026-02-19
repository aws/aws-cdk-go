//go:build no_runtime_type_checking

package previewawsm2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnApplicationConfigLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationConfigLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationConfigLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationConfigLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

