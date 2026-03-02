//go:build no_runtime_type_checking

package previewawssesmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMailManagerIngressPointApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnMailManagerIngressPointApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnMailManagerIngressPointApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnMailManagerIngressPointApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnMailManagerIngressPointApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnMailManagerIngressPointApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnMailManagerIngressPointApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnMailManagerIngressPointApplicationLogsS3Props) error {
	return nil
}

