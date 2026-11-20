//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnResourceMetadataContext) validateSetParameters(context *CfnResourceMetadataContextProps) error {
	return nil
}

func validateCfnResourceMetadataContext_OfParameters(scope constructs.IConstruct, options *CfnResourceMetadataContextOptions) error {
	return nil
}

