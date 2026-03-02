//go:build no_runtime_type_checking

package previewawsbedrockagentcoremixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWorkloadIdentityApplicationLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnWorkloadIdentityApplicationLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnWorkloadIdentityApplicationLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnWorkloadIdentityApplicationLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnWorkloadIdentityApplicationLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnWorkloadIdentityApplicationLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnWorkloadIdentityApplicationLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnWorkloadIdentityApplicationLogsS3Props) error {
	return nil
}

