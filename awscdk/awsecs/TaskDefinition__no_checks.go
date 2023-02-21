//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskDefinition) validateAddContainerParameters(id *string, props *ContainerDefinitionOptions) error {
	return nil
}

func (t *jsiiProxy_TaskDefinition) validateAddExtensionParameters(extension ITaskDefinitionExtension) error {
	return nil
}

func (t *jsiiProxy_TaskDefinition) validateAddFirelensLogRouterParameters(id *string, props *FirelensLogRouterDefinitionOptions) error {
	return nil
}

func (t *jsiiProxy_TaskDefinition) validateAddInferenceAcceleratorParameters(inferenceAccelerator *InferenceAccelerator) error {
	return nil
}

func (t *jsiiProxy_TaskDefinition) validateAddPlacementConstraintParameters(constraint PlacementConstraint) error {
	return nil
}

func (t *jsiiProxy_TaskDefinition) validateAddToExecutionRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (t *jsiiProxy_TaskDefinition) validateAddToTaskRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (t *jsiiProxy_TaskDefinition) validateAddVolumeParameters(volume *Volume) error {
	return nil
}

func (t *jsiiProxy_TaskDefinition) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TaskDefinition) validateFindContainerParameters(containerName *string) error {
	return nil
}

func (t *jsiiProxy_TaskDefinition) validateFindPortMappingByNameParameters(name *string) error {
	return nil
}

func (t *jsiiProxy_TaskDefinition) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TaskDefinition) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_TaskDefinition) validateGrantRunParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateTaskDefinition_FromTaskDefinitionArnParameters(scope constructs.Construct, id *string, taskDefinitionArn *string) error {
	return nil
}

func validateTaskDefinition_FromTaskDefinitionAttributesParameters(scope constructs.Construct, id *string, attrs *TaskDefinitionAttributes) error {
	return nil
}

func validateTaskDefinition_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTaskDefinition_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTaskDefinition_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTaskDefinitionParameters(scope constructs.Construct, id *string, props *TaskDefinitionProps) error {
	return nil
}

