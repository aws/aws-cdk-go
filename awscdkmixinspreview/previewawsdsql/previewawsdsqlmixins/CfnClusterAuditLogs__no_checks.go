//go:build no_runtime_type_checking

package previewawsdsqlmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClusterAuditLogs) validateToDestinationParameters(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnClusterAuditLogsDestProps) error {
	return nil
}

func (c *jsiiProxy_CfnClusterAuditLogs) validateToS3Parameters(bucket interfacesawss3.IBucketRef, props *CfnClusterAuditLogsS3Props) error {
	return nil
}

