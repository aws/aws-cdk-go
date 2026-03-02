//go:build no_runtime_type_checking

package previewawsmskmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnReplicatorApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnReplicatorApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnReplicatorApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnReplicatorApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnReplicatorApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnReplicatorApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnReplicatorApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnReplicatorApplicationLogsS3Props) error {
	return nil
}

