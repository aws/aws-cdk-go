//go:build no_runtime_type_checking

package previewawsnetworkfirewallmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFirewallAlertLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallAlertLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnFirewallAlertLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallAlertLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnFirewallAlertLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallAlertLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnFirewallAlertLogsS3Props) error {
	return nil
}

