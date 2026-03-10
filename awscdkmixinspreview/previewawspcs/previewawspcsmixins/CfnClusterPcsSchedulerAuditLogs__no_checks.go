//go:build no_runtime_type_checking

package previewawspcsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterPcsSchedulerAuditLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnClusterPcsSchedulerAuditLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterPcsSchedulerAuditLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnClusterPcsSchedulerAuditLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterPcsSchedulerAuditLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnClusterPcsSchedulerAuditLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterPcsSchedulerAuditLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnClusterPcsSchedulerAuditLogsS3Props) error {
	return nil
}

