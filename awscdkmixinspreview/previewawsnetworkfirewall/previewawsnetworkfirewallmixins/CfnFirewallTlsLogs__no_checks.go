//go:build no_runtime_type_checking

package previewawsnetworkfirewallmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFirewallTlsLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallTlsLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnFirewallTlsLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallTlsLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnFirewallTlsLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallTlsLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnFirewallTlsLogsS3Props) error {
	return nil
}

