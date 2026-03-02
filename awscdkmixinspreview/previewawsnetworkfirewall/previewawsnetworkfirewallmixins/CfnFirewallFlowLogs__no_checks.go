//go:build no_runtime_type_checking

package previewawsnetworkfirewallmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFirewallFlowLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallFlowLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnFirewallFlowLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallFlowLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnFirewallFlowLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallFlowLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnFirewallFlowLogsS3Props) error {
	return nil
}

