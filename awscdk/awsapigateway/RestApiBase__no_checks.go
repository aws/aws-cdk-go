//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RestApiBase) validateAddApiKeyParameters(id *string, options *ApiKeyOptions) error {
	return nil
}

func (r *jsiiProxy_RestApiBase) validateAddDomainNameParameters(id *string, options *DomainNameOptions) error {
	return nil
}

func (r *jsiiProxy_RestApiBase) validateAddGatewayResponseParameters(id *string, options *GatewayResponseOptions) error {
	return nil
}

func (r *jsiiProxy_RestApiBase) validateAddUsagePlanParameters(id *string, props *UsagePlanProps) error {
	return nil
}

func (r *jsiiProxy_RestApiBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RestApiBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RestApiBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_RestApiBase) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RestApiBase) validateMetricCacheHitCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RestApiBase) validateMetricCacheMissCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RestApiBase) validateMetricClientErrorParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RestApiBase) validateMetricCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RestApiBase) validateMetricIntegrationLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RestApiBase) validateMetricLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RestApiBase) validateMetricServerErrorParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateRestApiBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRestApiBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRestApiBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_RestApiBase) validateSetDeploymentStageParameters(val Stage) error {
	return nil
}

func validateNewRestApiBaseParameters(scope constructs.Construct, id *string, props *RestApiBaseProps) error {
	return nil
}

