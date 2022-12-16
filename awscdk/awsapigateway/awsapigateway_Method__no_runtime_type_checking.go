//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Method) validateAddMethodResponseParameters(methodResponse *MethodResponse) error {
	return nil
}

func (m *jsiiProxy_Method) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (m *jsiiProxy_Method) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (m *jsiiProxy_Method) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricParameters(metricName *string, stage IStage, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricCacheHitCountParameters(stage IStage, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricCacheMissCountParameters(stage IStage, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricClientErrorParameters(stage IStage, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricCountParameters(stage IStage, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricIntegrationLatencyParameters(stage IStage, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricLatencyParameters(stage IStage, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricServerErrorParameters(stage IStage, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateMethod_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMethod_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateMethod_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewMethodParameters(scope constructs.Construct, id *string, props *MethodProps) error {
	return nil
}

