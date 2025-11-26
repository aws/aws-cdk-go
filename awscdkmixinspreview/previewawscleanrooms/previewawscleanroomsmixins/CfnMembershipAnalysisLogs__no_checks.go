//go:build no_runtime_type_checking

package previewawscleanroomsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMembershipAnalysisLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnMembershipAnalysisLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnMembershipAnalysisLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

