//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VPNGateway) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VPNGateway) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VPNGateway) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateVPNGateway_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVPNGateway_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVPNGateway_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVPNGatewayParameters(scope constructs.Construct, id *string, props *VPNGatewayProps) error {
	return nil
}

