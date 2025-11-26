//go:build no_runtime_type_checking

package previewawsmskmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterBrokerLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnClusterBrokerLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnClusterBrokerLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

