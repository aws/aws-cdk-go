//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnOutput) validateOverrideLogicalIdParameters(newLogicalId *string) error {
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

