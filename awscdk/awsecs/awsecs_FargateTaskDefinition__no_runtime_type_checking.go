//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FargateTaskDefinition) validateAddContainerParameters(id *string, props *ContainerDefinitionOptions) error {
	return nil
}

func (f *jsiiProxy_FargateTaskDefinition) validateAddExtensionParameters(extension ITaskDefinitionExtension) error {
	return nil
}

func (f *jsiiProxy_FargateTaskDefinition) validateAddFirelensLogRouterParameters(id *string, props *FirelensLogRouterDefinitionOptions) error {
	return nil
}

func (f *jsiiProxy_FargateTaskDefinition) validateAddInferenceAcceleratorParameters(inferenceAccelerator *InferenceAccelerator) error {
	return nil
}

func (f *jsiiProxy_FargateTaskDefinition) validateAddPlacementConstraintParameters(constraint PlacementConstraint) error {
	return nil
}

func (f *jsiiProxy_FargateTaskDefinition) validateAddToExecutionRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (f *jsiiProxy_FargateTaskDefinition) validateAddToTaskRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (f *jsiiProxy_FargateTaskDefinition) validateAddVolumeParameters(volume *Volume) error {
	return nil
}

func (f *jsiiProxy_FargateTaskDefinition) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FargateTaskDefinition) validateFindContainerParameters(containerName *string) error {
	return nil
}

func (f *jsiiProxy_FargateTaskDefinition) validateFindPortMappingByNameParameters(name *string) error {
	return nil
}

func (f *jsiiProxy_FargateTaskDefinition) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FargateTaskDefinition) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (f *jsiiProxy_FargateTaskDefinition) validateGrantRunParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateFargateTaskDefinition_FromFargateTaskDefinitionArnParameters(scope constructs.Construct, id *string, fargateTaskDefinitionArn *string) error {
	return nil
}

func validateFargateTaskDefinition_FromFargateTaskDefinitionAttributesParameters(scope constructs.Construct, id *string, attrs *FargateTaskDefinitionAttributes) error {
	return nil
}

func validateFargateTaskDefinition_FromTaskDefinitionArnParameters(scope constructs.Construct, id *string, taskDefinitionArn *string) error {
	return nil
}

func validateFargateTaskDefinition_FromTaskDefinitionAttributesParameters(scope constructs.Construct, id *string, attrs *TaskDefinitionAttributes) error {
	return nil
}

func validateFargateTaskDefinition_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFargateTaskDefinition_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFargateTaskDefinition_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFargateTaskDefinitionParameters(scope constructs.Construct, id *string, props *FargateTaskDefinitionProps) error {
	return nil
}

