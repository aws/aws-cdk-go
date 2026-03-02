//go:build no_runtime_type_checking

package previewawsapsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWorkspaceManagedPrometheusLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnWorkspaceManagedPrometheusLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnWorkspaceManagedPrometheusLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnWorkspaceManagedPrometheusLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnWorkspaceManagedPrometheusLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnWorkspaceManagedPrometheusLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnWorkspaceManagedPrometheusLogsS3Props) error {
	return nil
}

