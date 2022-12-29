//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StageBase) validateAddApiKeyParameters(id *string, options *ApiKeyOptions) error {
	return nil
}

func (s *jsiiProxy_StageBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_StageBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_StageBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_StageBase) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StageBase) validateMetricCacheHitCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StageBase) validateMetricCacheMissCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StageBase) validateMetricClientErrorParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StageBase) validateMetricCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StageBase) validateMetricIntegrationLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StageBase) validateMetricLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StageBase) validateMetricServerErrorParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateStageBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStageBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateStageBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewStageBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

