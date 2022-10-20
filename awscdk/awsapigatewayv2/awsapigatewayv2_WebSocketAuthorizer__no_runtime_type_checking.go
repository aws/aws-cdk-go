//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapigatewayv2

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebSocketAuthorizer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WebSocketAuthorizer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WebSocketAuthorizer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (w *jsiiProxy_WebSocketAuthorizer) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (w *jsiiProxy_WebSocketAuthorizer) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateWebSocketAuthorizer_FromWebSocketAuthorizerAttributesParameters(scope constructs.Construct, id *string, attrs *WebSocketAuthorizerAttributes) error {
	return nil
}

func validateWebSocketAuthorizer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWebSocketAuthorizer_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewWebSocketAuthorizerParameters(scope constructs.Construct, id *string, props *WebSocketAuthorizerProps) error {
	return nil
}

