//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsDeploymentGroup) validateAddAlarmParameters(alarm awscloudwatch.IAlarm) error {
	return nil
}

func (e *jsiiProxy_EcsDeploymentGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EcsDeploymentGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EcsDeploymentGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateEcsDeploymentGroup_FromEcsDeploymentGroupAttributesParameters(scope constructs.Construct, id *string, attrs *EcsDeploymentGroupAttributes) error {
	return nil
}

func validateEcsDeploymentGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEcsDeploymentGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEcsDeploymentGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEcsDeploymentGroupParameters(scope constructs.Construct, id *string, props *EcsDeploymentGroupProps) error {
	return nil
}

