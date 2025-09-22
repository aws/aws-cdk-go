//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomResourceProviderBase) validateAddToRolePolicyParameters(statement interface{}) error {
	return nil
}

func validateCustomResourceProviderBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCustomResourceProviderBaseParameters(scope constructs.Construct, id *string, props *CustomResourceProviderBaseProps) error {
	return nil
}

