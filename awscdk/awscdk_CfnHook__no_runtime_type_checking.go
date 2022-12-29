//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHook) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateCfnHook_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnHook_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnHookParameters(scope constructs.Construct, id *string, props *CfnHookProps) error {
	return nil
}

