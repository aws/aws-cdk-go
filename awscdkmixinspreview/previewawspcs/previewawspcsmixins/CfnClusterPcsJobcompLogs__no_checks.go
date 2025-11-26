//go:build no_runtime_type_checking

package previewawspcsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterPcsJobcompLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnClusterPcsJobcompLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnClusterPcsJobcompLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

