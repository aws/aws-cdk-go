//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsDeploymentConfig) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EcsDeploymentConfig) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EcsDeploymentConfig) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateEcsDeploymentConfig_FromDeploymentConfigNameParameters(scope constructs.Construct, id *string, deploymentConfigName *string) error {
	return nil
}

func validateEcsDeploymentConfig_FromEcsDeploymentConfigNameParameters(scope constructs.Construct, id *string, ecsDeploymentConfigName *string) error {
	return nil
}

func validateEcsDeploymentConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEcsDeploymentConfig_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEcsDeploymentConfig_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEcsDeploymentConfigParameters(scope constructs.Construct, id *string, props *EcsDeploymentConfigProps) error {
	return nil
}

