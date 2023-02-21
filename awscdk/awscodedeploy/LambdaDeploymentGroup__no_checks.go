//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaDeploymentGroup) validateAddAlarmParameters(alarm awscloudwatch.IAlarm) error {
	return nil
}

func (l *jsiiProxy_LambdaDeploymentGroup) validateAddPostHookParameters(postHook awslambda.IFunction) error {
	return nil
}

func (l *jsiiProxy_LambdaDeploymentGroup) validateAddPreHookParameters(preHook awslambda.IFunction) error {
	return nil
}

func (l *jsiiProxy_LambdaDeploymentGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_LambdaDeploymentGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_LambdaDeploymentGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (l *jsiiProxy_LambdaDeploymentGroup) validateGrantPutLifecycleEventHookExecutionStatusParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateLambdaDeploymentGroup_FromLambdaDeploymentGroupAttributesParameters(scope constructs.Construct, id *string, attrs *LambdaDeploymentGroupAttributes) error {
	return nil
}

func validateLambdaDeploymentGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLambdaDeploymentGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateLambdaDeploymentGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewLambdaDeploymentGroupParameters(scope constructs.Construct, id *string, props *LambdaDeploymentGroupProps) error {
	return nil
}

