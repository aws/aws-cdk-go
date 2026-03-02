//go:build no_runtime_type_checking

package previewawsbackupgatewaymixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHypervisorDataAccessLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnHypervisorDataAccessLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnHypervisorDataAccessLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnHypervisorDataAccessLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnHypervisorDataAccessLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnHypervisorDataAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnHypervisorDataAccessLogsS3Props) error {
	return nil
}

