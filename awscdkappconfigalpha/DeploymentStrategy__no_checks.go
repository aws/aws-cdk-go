//go:build no_runtime_type_checking

package awscdkappconfigalpha

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeploymentStrategy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DeploymentStrategy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DeploymentStrategy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateDeploymentStrategy_FromDeploymentStrategyArnParameters(scope constructs.Construct, id *string, deploymentStrategyArn *string) error {
	return nil
}

func validateDeploymentStrategy_FromDeploymentStrategyIdParameters(scope constructs.Construct, id *string, deploymentStrategyId DeploymentStrategyId) error {
	return nil
}

func validateDeploymentStrategy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDeploymentStrategy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDeploymentStrategy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDeploymentStrategyParameters(scope constructs.Construct, id *string, props *DeploymentStrategyProps) error {
	return nil
}

