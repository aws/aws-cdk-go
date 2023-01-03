//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BaseDeploymentConfig) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BaseDeploymentConfig) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BaseDeploymentConfig) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateBaseDeploymentConfig_FromDeploymentConfigNameParameters(scope constructs.Construct, id *string, deploymentConfigName *string) error {
	return nil
}

func validateBaseDeploymentConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBaseDeploymentConfig_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBaseDeploymentConfig_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBaseDeploymentConfigParameters(scope constructs.Construct, id *string, props *BaseDeploymentConfigProps) error {
	return nil
}

