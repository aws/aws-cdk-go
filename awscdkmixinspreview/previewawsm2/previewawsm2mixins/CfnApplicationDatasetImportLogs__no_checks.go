//go:build no_runtime_type_checking

package previewawsm2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnApplicationDatasetImportLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationDatasetImportLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnApplicationDatasetImportLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationDatasetImportLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnApplicationDatasetImportLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnApplicationDatasetImportLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnApplicationDatasetImportLogsS3Props) error {
	return nil
}

