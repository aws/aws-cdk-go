//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnParameter) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnParameter) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnParameter) validateResolveParameters(_context IResolveContext) error {
	return nil
}

func (c *jsiiProxy_CfnParameter) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func validateCfnParameter_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnParameter_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnParameter) validateSetDefaultParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnParameter) validateSetNoEchoParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_CfnParameter) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewCfnParameterParameters(scope constructs.Construct, id *string, props *CfnParameterProps) error {
	return nil
}

