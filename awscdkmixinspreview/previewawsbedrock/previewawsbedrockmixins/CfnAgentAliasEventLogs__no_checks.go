//go:build no_runtime_type_checking

package previewawsbedrockmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAgentAliasEventLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnAgentAliasEventLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnAgentAliasEventLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

