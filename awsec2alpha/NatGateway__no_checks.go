//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NatGateway) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NatGateway) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NatGateway) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateNatGateway_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNatGateway_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNatGateway_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNatGatewayParameters(scope constructs.Construct, id *string, props *NatGatewayProps) error {
	return nil
}

