//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpAuthorizer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (h *jsiiProxy_HttpAuthorizer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (h *jsiiProxy_HttpAuthorizer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateHttpAuthorizer_FromHttpAuthorizerAttributesParameters(scope constructs.Construct, id *string, attrs *HttpAuthorizerAttributes) error {
	return nil
}

func validateHttpAuthorizer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHttpAuthorizer_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateHttpAuthorizer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewHttpAuthorizerParameters(scope constructs.Construct, id *string, props *HttpAuthorizerProps) error {
	return nil
}

