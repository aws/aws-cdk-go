//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsApplication) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EcsApplication) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EcsApplication) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateEcsApplication_FromEcsApplicationNameParameters(scope constructs.Construct, id *string, ecsApplicationName *string) error {
	return nil
}

func validateEcsApplication_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEcsApplication_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEcsApplication_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEcsApplicationParameters(scope constructs.Construct, id *string, props *EcsApplicationProps) error {
	return nil
}

