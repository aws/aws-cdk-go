//go:build no_runtime_type_checking

package previewawsorganizationsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnOrganizationWorkmailPersonalAccessTokenLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnOrganizationWorkmailPersonalAccessTokenLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnOrganizationWorkmailPersonalAccessTokenLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

