//go:build no_runtime_type_checking

package previewawstransfermixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnServerTransferLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnServerTransferLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnServerTransferLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnServerTransferLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnServerTransferLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnServerTransferLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnServerTransferLogsS3Props) error {
	return nil
}

