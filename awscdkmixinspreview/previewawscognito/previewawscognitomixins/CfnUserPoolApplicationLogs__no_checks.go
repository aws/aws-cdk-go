//go:build no_runtime_type_checking

package previewawscognitomixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnUserPoolApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnUserPoolApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnUserPoolApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

