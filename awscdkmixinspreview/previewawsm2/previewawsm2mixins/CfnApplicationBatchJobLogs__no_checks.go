//go:build no_runtime_type_checking

package previewawsm2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnApplicationBatchJobLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationBatchJobLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnApplicationBatchJobLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationBatchJobLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnApplicationBatchJobLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationBatchJobLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnApplicationBatchJobLogsS3Props) error {
	return nil
}

