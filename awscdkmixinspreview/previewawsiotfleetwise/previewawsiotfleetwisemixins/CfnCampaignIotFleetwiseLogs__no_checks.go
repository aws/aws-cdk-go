//go:build no_runtime_type_checking

package previewawsiotfleetwisemixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCampaignIotFleetwiseLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnCampaignIotFleetwiseLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef, props *CfnCampaignIotFleetwiseLogsFirehoseProps) error {
	return nil
}

func (c *jsiiProxy_CfnCampaignIotFleetwiseLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef, props *CfnCampaignIotFleetwiseLogsLogGroupProps) error {
	return nil
}

func (c *jsiiProxy_CfnCampaignIotFleetwiseLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnCampaignIotFleetwiseLogsS3Props) error {
	return nil
}

