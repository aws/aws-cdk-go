//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapigatewayv2

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebSocketRoute) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WebSocketRoute) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WebSocketRoute) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (w *jsiiProxy_WebSocketRoute) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (w *jsiiProxy_WebSocketRoute) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateWebSocketRoute_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWebSocketRoute_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewWebSocketRouteParameters(scope constructs.Construct, id *string, props *WebSocketRouteProps) error {
	return nil
}

