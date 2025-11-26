//go:build no_runtime_type_checking

package previewawsqbusinessmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnApplicationSyncJobLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationSyncJobLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationSyncJobLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

