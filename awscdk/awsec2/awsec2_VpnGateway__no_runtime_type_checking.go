//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpnGateway) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VpnGateway) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VpnGateway) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (v *jsiiProxy_VpnGateway) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (v *jsiiProxy_VpnGateway) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateVpnGateway_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVpnGateway_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewVpnGatewayParameters(scope constructs.Construct, id *string, props *VpnGatewayProps) error {
	return nil
}

