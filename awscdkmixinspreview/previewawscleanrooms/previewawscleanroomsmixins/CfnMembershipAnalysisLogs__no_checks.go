//go:build no_runtime_type_checking

package previewawscleanroomsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMembershipAnalysisLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnMembershipAnalysisLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnMembershipAnalysisLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnMembershipAnalysisLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnMembershipAnalysisLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnMembershipAnalysisLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnMembershipAnalysisLogsS3Props) error {
	return nil
}

