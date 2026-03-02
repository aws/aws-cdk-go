//go:build no_runtime_type_checking

package previewawswafv2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWebACLAccessLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnWebACLAccessLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnWebACLAccessLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnWebACLAccessLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnWebACLAccessLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnWebACLAccessLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnWebACLAccessLogsS3Props) error {
	return nil
}

