//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServerDeploymentGroup) validateAddAlarmParameters(alarm awscloudwatch.IAlarm) error {
	return nil
}

func (s *jsiiProxy_ServerDeploymentGroup) validateAddAutoScalingGroupParameters(asg awsautoscaling.AutoScalingGroup) error {
	return nil
}

func (s *jsiiProxy_ServerDeploymentGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_ServerDeploymentGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_ServerDeploymentGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateServerDeploymentGroup_FromServerDeploymentGroupAttributesParameters(scope constructs.Construct, id *string, attrs *ServerDeploymentGroupAttributes) error {
	return nil
}

func validateServerDeploymentGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateServerDeploymentGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateServerDeploymentGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewServerDeploymentGroupParameters(scope constructs.Construct, id *string, props *ServerDeploymentGroupProps) error {
	return nil
}

