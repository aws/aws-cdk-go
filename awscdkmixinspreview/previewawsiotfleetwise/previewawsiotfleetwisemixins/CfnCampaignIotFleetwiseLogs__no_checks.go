//go:build no_runtime_type_checking

package previewawsiotfleetwisemixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCampaignIotFleetwiseLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnCampaignIotFleetwiseLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnCampaignIotFleetwiseLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

