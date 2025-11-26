//go:build no_runtime_type_checking

package previewawswafv2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWebACLAccessLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnWebACLAccessLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnWebACLAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

