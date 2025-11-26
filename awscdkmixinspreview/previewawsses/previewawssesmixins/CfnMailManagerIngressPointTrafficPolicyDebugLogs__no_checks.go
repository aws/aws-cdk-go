//go:build no_runtime_type_checking

package previewawssesmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

