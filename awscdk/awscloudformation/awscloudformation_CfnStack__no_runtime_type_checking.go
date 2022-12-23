//go:build no_runtime_type_checking

package awscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStack) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateAddDependencyParameters(target awscdk.CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateGetMetadataParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateInspectParameters(inspector awscdk.TreeInspector) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateRemoveDependencyParameters(target awscdk.CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateReplaceDependencyParameters(target awscdk.CfnResource, newTarget awscdk.CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnStack) validateValidatePropertiesParameters(_properties interface{}) error {
	return nil
}

func validateCfnStack_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnStack_IsCfnResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnStack_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnStack) validateSetParametersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnStack) validateSetTemplateUrlParameters(val *string) error {
	return nil
}

func validateNewCfnStackParameters(scope constructs.Construct, id *string, props *CfnStackProps) error {
	return nil
}

