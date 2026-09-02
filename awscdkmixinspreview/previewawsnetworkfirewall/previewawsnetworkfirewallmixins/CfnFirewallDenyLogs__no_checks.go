//go:build no_runtime_type_checking

package previewawsnetworkfirewallmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFirewallDenyLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnFirewallDenyLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallDenyLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnFirewallDenyLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallDenyLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnFirewallDenyLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallDenyLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnFirewallDenyLogsS3Props) error {
	return nil
}

