//go:build no_runtime_type_checking

package previewawsbackupgatewaymixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHypervisorDataAccessLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnHypervisorDataAccessLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnHypervisorDataAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

