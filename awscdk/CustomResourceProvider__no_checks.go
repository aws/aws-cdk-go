//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomResourceProvider) validateAddToRolePolicyParameters(statement interface{}) error {
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

