//go:build no_runtime_type_checking

package previewawsnetworkfirewallmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFirewallTlsLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallTlsLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallTlsLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallTlsLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

