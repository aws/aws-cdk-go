//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnJson) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnJson) validateResolveParameters(_arg IResolveContext) error {
	return nil
}

func (c *jsiiProxy_CfnJson) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func validateCfnJson_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnJsonParameters(scope constructs.Construct, id *string, props *CfnJsonProps) error {
	return nil
}

