//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRefElement) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateCfnRefElement_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnRefElement_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnRefElementParameters(scope constructs.Construct, id *string) error {
	return nil
}

