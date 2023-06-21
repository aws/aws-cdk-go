//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServerDeploymentConfig) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_ServerDeploymentConfig) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_ServerDeploymentConfig) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateServerDeploymentConfig_FromDeploymentConfigNameParameters(scope constructs.Construct, id *string, deploymentConfigName *string) error {
	return nil
}

func validateServerDeploymentConfig_FromServerDeploymentConfigNameParameters(scope constructs.Construct, id *string, serverDeploymentConfigName *string) error {
	return nil
}

func validateServerDeploymentConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateServerDeploymentConfig_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateServerDeploymentConfig_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewServerDeploymentConfigParameters(scope constructs.Construct, id *string, props *ServerDeploymentConfigProps) error {
	return nil
}

