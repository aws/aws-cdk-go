//go:build no_runtime_type_checking

package previewawsbackupgatewaymixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHypervisorBgwHypervisorLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnHypervisorBgwHypervisorLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnHypervisorBgwHypervisorLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnHypervisorBgwHypervisorLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

