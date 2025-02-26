//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Deployment) validateAddToLogicalIdParameters(data interface{}) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateDeployment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDeployment_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDeployment_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDeploymentParameters(scope constructs.Construct, id *string, props *DeploymentProps) error {
	return nil
}

