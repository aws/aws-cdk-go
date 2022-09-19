//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

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

func validateHttpRoute_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHttpRoute_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateHttpRoute_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewHttpRouteParameters(scope constructs.Construct, id *string, props *HttpRouteProps) error {
	return nil
}

