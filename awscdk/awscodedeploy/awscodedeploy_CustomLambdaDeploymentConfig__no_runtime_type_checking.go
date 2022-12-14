//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomLambdaDeploymentConfig) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CustomLambdaDeploymentConfig) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CustomLambdaDeploymentConfig) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCustomLambdaDeploymentConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCustomLambdaDeploymentConfig_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCustomLambdaDeploymentConfig_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCustomLambdaDeploymentConfigParameters(scope constructs.Construct, id *string, props *CustomLambdaDeploymentConfigProps) error {
	return nil
}

