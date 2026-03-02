//go:build no_runtime_type_checking

package previewawsec2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVPNConnectionEventLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnVPNConnectionEventLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnVPNConnectionEventLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnVPNConnectionEventLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnVPNConnectionEventLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnVPNConnectionEventLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnVPNConnectionEventLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnVPNConnectionEventLogsS3Props) error {
	return nil
}

