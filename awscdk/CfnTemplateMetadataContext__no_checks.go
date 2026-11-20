//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTemplateMetadataContext) validateSetParameters(context *CfnTemplateMetadataContextProps) error {
	return nil
}

func validateCfnTemplateMetadataContext_OfParameters(stack Stack) error {
	return nil
}

