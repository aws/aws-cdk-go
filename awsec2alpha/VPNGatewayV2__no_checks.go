//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VPNGatewayV2) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VPNGatewayV2) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VPNGatewayV2) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateVPNGatewayV2_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVPNGatewayV2_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVPNGatewayV2_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVPNGatewayV2Parameters(scope constructs.Construct, id *string, props *VPNGatewayV2Props) error {
	return nil
}

