//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_INetworkTargetGroupMetrics) validateCustomParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_INetworkTargetGroupMetrics) validateHealthyHostCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_INetworkTargetGroupMetrics) validateUnHealthyHostCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

