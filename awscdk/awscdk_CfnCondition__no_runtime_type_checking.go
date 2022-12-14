//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCondition) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnCondition) validateResolveParameters(_context IResolveContext) error {
	return nil
}

func validateCfnCondition_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnCondition_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnConditionParameters(scope constructs.Construct, id *string, props *CfnConditionProps) error {
	return nil
}

