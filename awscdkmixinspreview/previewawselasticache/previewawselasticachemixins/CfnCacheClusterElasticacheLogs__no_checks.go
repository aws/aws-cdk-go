//go:build no_runtime_type_checking

package previewawselasticachemixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCacheClusterElasticacheLogs) validateToFirehoseParameters(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

func (c *jsiiProxy_CfnCacheClusterElasticacheLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

func (c *jsiiProxy_CfnCacheClusterElasticacheLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

