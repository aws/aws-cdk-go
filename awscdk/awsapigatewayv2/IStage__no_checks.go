//go:build no_runtime_type_checking

package awsapigatewayv2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IStage) validateAddStageVariableParameters(name *string, value *string) error {
	return nil
}

func (i *jsiiProxy_IStage) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

