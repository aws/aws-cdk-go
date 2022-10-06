//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GatewayResponse) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_GatewayResponse) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GatewayResponse) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (g *jsiiProxy_GatewayResponse) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (g *jsiiProxy_GatewayResponse) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateGatewayResponse_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGatewayResponse_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewGatewayResponseParameters(scope constructs.Construct, id *string, props *GatewayResponseProps) error {
	return nil
}

