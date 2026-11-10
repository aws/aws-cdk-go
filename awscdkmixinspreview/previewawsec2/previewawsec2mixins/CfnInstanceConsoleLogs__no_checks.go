//go:build no_runtime_type_checking

package previewawsec2mixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInstanceConsoleLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnInstanceConsoleLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnInstanceConsoleLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnInstanceConsoleLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnInstanceConsoleLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnInstanceConsoleLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnInstanceConsoleLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnInstanceConsoleLogsS3Props) error {
	return nil
}

