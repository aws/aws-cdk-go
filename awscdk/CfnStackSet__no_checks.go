//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStackSet) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateAddDependencyParameters(target CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateAddDependsOnParameters(target CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateApplyRemovalPolicyParameters(options *RemovalPolicyOptions) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateGetMetadataParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateInspectParameters(inspector TreeInspector) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateRemoveDependencyParameters(target CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateReplaceDependencyParameters(target CfnResource, newTarget CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnStackSet) validateValidatePropertiesParameters(_properties interface{}) error {
	return nil
}

func validateCfnStackSet_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnStackSet_IsCfnResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnStackSet_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnStackSet) validateSetAutoDeploymentParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnStackSet) validateSetManagedExecutionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnStackSet) validateSetOperationPreferencesParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnStackSet) validateSetParametersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnStackSet) validateSetPermissionModelParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CfnStackSet) validateSetStackInstancesGroupParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnStackSet) validateSetStackSetNameParameters(val *string) error {
	return nil
}

func validateNewCfnStackSetParameters(scope constructs.Construct, id *string, props *CfnStackSetProps) error {
	return nil
}

