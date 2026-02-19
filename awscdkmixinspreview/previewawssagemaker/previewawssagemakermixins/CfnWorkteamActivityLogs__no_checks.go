//go:build no_runtime_type_checking

package previewawssagemakermixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWorkteamActivityLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef) error {
	return nil
}

func (c *jsiiProxy_CfnWorkteamActivityLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnWorkteamActivityLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

