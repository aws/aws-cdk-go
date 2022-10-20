//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapigatewayv2

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpRoute) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (h *jsiiProxy_HttpRoute) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (h *jsiiProxy_HttpRoute) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (h *jsiiProxy_HttpRoute) validateGrantInvokeParameters(grantee awsiam.IGrantable, options *GrantInvokeOptions) error {
	return nil
}

func (h *jsiiProxy_HttpRoute) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (h *jsiiProxy_HttpRoute) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateHttpRoute_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHttpRoute_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewHttpRouteParameters(scope constructs.Construct, id *string, props *HttpRouteProps) error {
	return nil
}

