//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInclude) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnInclude) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnInclude) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func validateCfnInclude_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnInclude_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnIncludeParameters(scope constructs.Construct, id *string, props *CfnIncludeProps) error {
	return nil
}

