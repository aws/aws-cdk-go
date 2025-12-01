//go:build no_runtime_type_checking

package previewawsrummixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAppMonitorRumOtelSpans) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnAppMonitorRumOtelSpans) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnAppMonitorRumOtelSpans) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

