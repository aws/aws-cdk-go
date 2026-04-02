//go:build no_runtime_type_checking

package awsmediapackagev2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IChannel) validateAddOriginEndpointParameters(id *string, options *OriginEndpointOptions) error {
	return nil
}

func (i *jsiiProxy_IChannel) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IChannel) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IChannel) validateMetricEgressBytesParameters(options *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IChannel) validateMetricEgressRequestCountParameters(options *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IChannel) validateMetricEgressResponseTimeParameters(options *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IChannel) validateMetricIngressBytesParameters(options *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IChannel) validateMetricIngressRequestCountParameters(options *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IChannel) validateMetricIngressResponseTimeParameters(options *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IChannel) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

