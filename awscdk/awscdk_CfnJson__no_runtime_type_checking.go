//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnJson) validateResolveParameters(_arg IResolveContext) error {
	return nil
}

func validateCfnJson_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnJsonParameters(scope constructs.Construct, id *string, props *CfnJsonProps) error {
	return nil
}

