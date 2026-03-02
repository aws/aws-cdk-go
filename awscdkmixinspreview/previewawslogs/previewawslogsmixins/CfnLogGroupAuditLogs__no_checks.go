//go:build no_runtime_type_checking

package previewawslogsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLogGroupAuditLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnLogGroupAuditLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroupAuditLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnLogGroupAuditLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroupAuditLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnLogGroupAuditLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroupAuditLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnLogGroupAuditLogsS3Props) error {
	return nil
}

