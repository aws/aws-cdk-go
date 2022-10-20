//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnOutput) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnOutput) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnOutput) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func validateCfnOutput_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnOutput_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnOutput) validateSetValueParameters(val interface{}) error {
	return nil
}

func validateNewCfnOutputParameters(scope constructs.Construct, id *string, props *CfnOutputProps) error {
	return nil
}

