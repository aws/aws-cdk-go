//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnResource) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_CfnResource) validateAddDependencyParameters(target CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnResource) validateAddDependsOnParameters(target CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnResource) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnResource) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnResource) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	return nil
}

func (c *jsiiProxy_CfnResource) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnResource) validateApplyRemovalPolicyParameters(options *RemovalPolicyOptions) error {
	return nil
}

func (c *jsiiProxy_CfnResource) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CfnResource) validateGetMetadataParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CfnResource) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnResource) validateRemoveDependencyParameters(target CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnResource) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnResource) validateReplaceDependencyParameters(target CfnResource, newTarget CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnResource) validateValidatePropertiesParameters(_properties interface{}) error {
	return nil
}

func validateCfnResource_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnResource_IsCfnResourceParameters(x interface{}) error {
	return nil
}

func validateCfnResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnResourceParameters(scope constructs.Construct, id *string, props *CfnResourceProps) error {
	return nil
}

