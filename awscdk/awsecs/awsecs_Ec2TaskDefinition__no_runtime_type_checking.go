//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2TaskDefinition) validateAddContainerParameters(id *string, props *ContainerDefinitionOptions) error {
	return nil
}

func (e *jsiiProxy_Ec2TaskDefinition) validateAddExtensionParameters(extension ITaskDefinitionExtension) error {
	return nil
}

func (e *jsiiProxy_Ec2TaskDefinition) validateAddFirelensLogRouterParameters(id *string, props *FirelensLogRouterDefinitionOptions) error {
	return nil
}

func (e *jsiiProxy_Ec2TaskDefinition) validateAddInferenceAcceleratorParameters(inferenceAccelerator *InferenceAccelerator) error {
	return nil
}

func (e *jsiiProxy_Ec2TaskDefinition) validateAddPlacementConstraintParameters(constraint PlacementConstraint) error {
	return nil
}

func (e *jsiiProxy_Ec2TaskDefinition) validateAddToExecutionRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (e *jsiiProxy_Ec2TaskDefinition) validateAddToTaskRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (e *jsiiProxy_Ec2TaskDefinition) validateAddVolumeParameters(volume *Volume) error {
	return nil
}

func (e *jsiiProxy_Ec2TaskDefinition) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_Ec2TaskDefinition) validateFindContainerParameters(containerName *string) error {
	return nil
}

func (e *jsiiProxy_Ec2TaskDefinition) validateFindPortMappingByNameParameters(name *string) error {
	return nil
}

func (e *jsiiProxy_Ec2TaskDefinition) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_Ec2TaskDefinition) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_Ec2TaskDefinition) validateGrantRunParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateEc2TaskDefinition_FromEc2TaskDefinitionArnParameters(scope constructs.Construct, id *string, ec2TaskDefinitionArn *string) error {
	return nil
}

func validateEc2TaskDefinition_FromEc2TaskDefinitionAttributesParameters(scope constructs.Construct, id *string, attrs *Ec2TaskDefinitionAttributes) error {
	return nil
}

func validateEc2TaskDefinition_FromTaskDefinitionArnParameters(scope constructs.Construct, id *string, taskDefinitionArn *string) error {
	return nil
}

func validateEc2TaskDefinition_FromTaskDefinitionAttributesParameters(scope constructs.Construct, id *string, attrs *TaskDefinitionAttributes) error {
	return nil
}

func validateEc2TaskDefinition_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEc2TaskDefinition_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEc2TaskDefinition_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEc2TaskDefinitionParameters(scope constructs.Construct, id *string, props *Ec2TaskDefinitionProps) error {
	return nil
}

