//go:build no_runtime_type_checking
// +build no_runtime_type_checking

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

func (c *jsiiProxy_CustomLambdaDeploymentConfig) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CustomLambdaDeploymentConfig) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateCustomLambdaDeploymentConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCustomLambdaDeploymentConfig_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewCustomLambdaDeploymentConfigParameters(scope constructs.Construct, id *string, props *CustomLambdaDeploymentConfigProps) error {
	return nil
}

