//go:build no_runtime_type_checking

package previewawsrummixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAppMonitorRumTelemetryLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnAppMonitorRumTelemetryLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnAppMonitorRumTelemetryLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnAppMonitorRumTelemetryLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnAppMonitorRumTelemetryLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnAppMonitorRumTelemetryLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnAppMonitorRumTelemetryLogsS3Props) error {
	return nil
}

