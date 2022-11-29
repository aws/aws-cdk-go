//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnModuleVersion) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersion) validateAddDependsOnParameters(target CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersion) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersion) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersion) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersion) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersion) validateApplyRemovalPolicyParameters(options *RemovalPolicyOptions) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersion) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersion) validateGetMetadataParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersion) validateInspectParameters(inspector TreeInspector) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersion) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersion) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersion) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersion) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnModuleVersion) validateValidatePropertiesParameters(_properties interface{}) error {
	return nil
}

func validateCfnModuleVersion_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnModuleVersion_IsCfnResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnModuleVersion_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnModuleVersion) validateSetModuleNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CfnModuleVersion) validateSetModulePackageParameters(val *string) error {
	return nil
}

func validateNewCfnModuleVersionParameters(scope Construct, id *string, props *CfnModuleVersionProps) error {
	return nil
}

