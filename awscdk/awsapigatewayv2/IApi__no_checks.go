//go:build no_runtime_type_checking

package awsapigatewayv2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IApi) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IApi) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

