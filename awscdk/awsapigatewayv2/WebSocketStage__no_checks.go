//go:build no_runtime_type_checking

package awsapigatewayv2

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebSocketStage) validateAddStageVariableParameters(name *string, value *string) error {
	return nil
}

func (w *jsiiProxy_WebSocketStage) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WebSocketStage) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WebSocketStage) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (w *jsiiProxy_WebSocketStage) validateGrantManagementApiAccessParameters(identity awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WebSocketStage) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateWebSocketStage_FromWebSocketStageAttributesParameters(scope constructs.Construct, id *string, attrs *WebSocketStageAttributes) error {
	return nil
}

func validateWebSocketStage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWebSocketStage_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWebSocketStage_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewWebSocketStageParameters(scope constructs.Construct, id *string, props *WebSocketStageProps) error {
	return nil
}

