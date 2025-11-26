//go:build no_runtime_type_checking

package previewawspipesmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPipeExecutionLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnPipeExecutionLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnPipeExecutionLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

