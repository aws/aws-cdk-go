//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomResourceProvider) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CustomResourceProvider) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func validateCustomResourceProvider_GetOrCreateParameters(scope constructs.Construct, uniqueid *string, props *CustomResourceProviderProps) error {
	return nil
}

func validateCustomResourceProvider_GetOrCreateProviderParameters(scope constructs.Construct, uniqueid *string, props *CustomResourceProviderProps) error {
	return nil
}

func validateCustomResourceProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCustomResourceProviderParameters(scope constructs.Construct, id *string, props *CustomResourceProviderProps) error {
	return nil
}

