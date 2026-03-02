//go:build no_runtime_type_checking

package previewawsmskmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterBrokerLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnClusterBrokerLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnClusterBrokerLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterBrokerLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnClusterBrokerLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterBrokerLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnClusterBrokerLogsS3Props) error {
	return nil
}

