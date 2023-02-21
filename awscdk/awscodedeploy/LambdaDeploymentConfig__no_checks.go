//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaDeploymentConfig) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_LambdaDeploymentConfig) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_LambdaDeploymentConfig) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateLambdaDeploymentConfig_FromDeploymentConfigNameParameters(scope constructs.Construct, id *string, deploymentConfigName *string) error {
	return nil
}

func validateLambdaDeploymentConfig_FromLambdaDeploymentConfigNameParameters(scope constructs.Construct, id *string, lambdaDeploymentConfigName *string) error {
	return nil
}

func validateLambdaDeploymentConfig_ImportParameters(_scope constructs.Construct, _id *string, props *LambdaDeploymentConfigImportProps) error {
	return nil
}

func validateLambdaDeploymentConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLambdaDeploymentConfig_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateLambdaDeploymentConfig_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewLambdaDeploymentConfigParameters(scope constructs.Construct, id *string, props *LambdaDeploymentConfigProps) error {
	return nil
}

