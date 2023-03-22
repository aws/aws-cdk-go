//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMapping) validateFindInMapParameters(key1 *string, key2 *string) error {
	return nil
}

func (c *jsiiProxy_CfnMapping) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnMapping) validateSetValueParameters(key1 *string, key2 *string, value interface{}) error {
	return nil
}

func validateCfnMapping_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnMapping_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnMappingParameters(scope constructs.Construct, id *string, props *CfnMappingProps) error {
	return nil
}

