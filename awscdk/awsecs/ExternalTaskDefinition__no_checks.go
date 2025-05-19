//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalTaskDefinition) validateAddContainerParameters(id *string, props *ContainerDefinitionOptions) error {
	return nil
}

func (e *jsiiProxy_ExternalTaskDefinition) validateAddExtensionParameters(extension ITaskDefinitionExtension) error {
	return nil
}

func (e *jsiiProxy_ExternalTaskDefinition) validateAddFirelensLogRouterParameters(id *string, props *FirelensLogRouterDefinitionOptions) error {
	return nil
}

func (e *jsiiProxy_ExternalTaskDefinition) validateAddInferenceAcceleratorParameters(_inferenceAccelerator *InferenceAccelerator) error {
	return nil
}

func (e *jsiiProxy_ExternalTaskDefinition) validateAddPlacementConstraintParameters(constraint PlacementConstraint) error {
	return nil
}

func (e *jsiiProxy_ExternalTaskDefinition) validateAddToExecutionRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (e *jsiiProxy_ExternalTaskDefinition) validateAddToTaskRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (e *jsiiProxy_ExternalTaskDefinition) validateAddVolumeParameters(volume *Volume) error {
	return nil
}

func (e *jsiiProxy_ExternalTaskDefinition) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_ExternalTaskDefinition) validateFindContainerParameters(containerName *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTaskDefinition) validateFindPortMappingByNameParameters(name *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTaskDefinition) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_ExternalTaskDefinition) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTaskDefinition) validateGrantRunParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateExternalTaskDefinition_FromEc2TaskDefinitionArnParameters(scope constructs.Construct, id *string, externalTaskDefinitionArn *string) error {
	return nil
}

func validateExternalTaskDefinition_FromExternalTaskDefinitionAttributesParameters(scope constructs.Construct, id *string, attrs *ExternalTaskDefinitionAttributes) error {
	return nil
}

func validateExternalTaskDefinition_FromTaskDefinitionArnParameters(scope constructs.Construct, id *string, taskDefinitionArn *string) error {
	return nil
}

func validateExternalTaskDefinition_FromTaskDefinitionAttributesParameters(scope constructs.Construct, id *string, attrs *TaskDefinitionAttributes) error {
	return nil
}

func validateExternalTaskDefinition_IsConstructParameters(x interface{}) error {
	return nil
}

func validateExternalTaskDefinition_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateExternalTaskDefinition_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewExternalTaskDefinitionParameters(scope constructs.Construct, id *string, props *ExternalTaskDefinitionProps) error {
	return nil
}

