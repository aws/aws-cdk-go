//go:build no_runtime_type_checking

package previewawsorganizationsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnOrganizationWorkmailAvailabilityProviderLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnOrganizationWorkmailAvailabilityProviderLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnOrganizationWorkmailAvailabilityProviderLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

