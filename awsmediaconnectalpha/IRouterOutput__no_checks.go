//go:build no_runtime_type_checking

package awsmediaconnectalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IRouterOutput) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IRouterOutput) validateMetricArqRequestsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IRouterOutput) validateMetricBitrateParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IRouterOutput) validateMetricConnectedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IRouterOutput) validateMetricTotalPacketsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IRouterOutput) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

