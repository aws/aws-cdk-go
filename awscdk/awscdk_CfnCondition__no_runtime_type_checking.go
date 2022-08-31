//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCondition) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnCondition) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnCondition) validateResolveParameters(_context IResolveContext) error {
	return nil
}

func (c *jsiiProxy_CfnCondition) validateSynthesizeParameters(session ISynthesisSession) error {
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

