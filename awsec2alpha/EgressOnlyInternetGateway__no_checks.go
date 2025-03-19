//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EgressOnlyInternetGateway) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EgressOnlyInternetGateway) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EgressOnlyInternetGateway) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateEgressOnlyInternetGateway_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEgressOnlyInternetGateway_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEgressOnlyInternetGateway_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEgressOnlyInternetGatewayParameters(scope constructs.Construct, id *string, props *EgressOnlyInternetGatewayProps) error {
	return nil
}

