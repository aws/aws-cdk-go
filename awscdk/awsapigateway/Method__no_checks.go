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

func (m *jsiiProxy_Method) validateGrantExecuteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricParameters(metricName *string, stage interfacesawsapigateway.IStageRef, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricCacheHitCountParameters(stage interfacesawsapigateway.IStageRef, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricCacheMissCountParameters(stage interfacesawsapigateway.IStageRef, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricClientErrorParameters(stage interfacesawsapigateway.IStageRef, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricCountParameters(stage interfacesawsapigateway.IStageRef, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricIntegrationLatencyParameters(stage interfacesawsapigateway.IStageRef, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricLatencyParameters(stage interfacesawsapigateway.IStageRef, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (m *jsiiProxy_Method) validateMetricServerErrorParameters(stage interfacesawsapigateway.IStageRef, props *awscloudwatch.MetricOptions) error {
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

