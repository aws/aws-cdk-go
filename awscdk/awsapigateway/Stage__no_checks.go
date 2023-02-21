//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Stage) validateAddApiKeyParameters(id *string, options *ApiKeyOptions) error {
	return nil
}

func (s *jsiiProxy_Stage) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_Stage) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Stage) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_Stage) validateMetricCacheHitCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_Stage) validateMetricCacheMissCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_Stage) validateMetricClientErrorParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_Stage) validateMetricCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_Stage) validateMetricIntegrationLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_Stage) validateMetricLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_Stage) validateMetricServerErrorParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateStage_FromStageAttributesParameters(scope constructs.Construct, id *string, attrs *StageAttributes) error {
	return nil
}

func validateStage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStage_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateStage_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewStageParameters(scope constructs.Construct, id *string, props *StageProps) error {
	return nil
}

