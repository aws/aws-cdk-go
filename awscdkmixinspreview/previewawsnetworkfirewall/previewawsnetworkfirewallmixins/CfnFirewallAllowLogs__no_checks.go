//go:build no_runtime_type_checking

package previewawsnetworkfirewallmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFirewallAllowLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnFirewallAllowLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallAllowLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnFirewallAllowLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallAllowLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnFirewallAllowLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallAllowLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnFirewallAllowLogsS3Props) error {
	return nil
}

