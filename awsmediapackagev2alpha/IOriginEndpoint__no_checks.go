//go:build no_runtime_type_checking

package awsmediapackagev2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IOriginEndpoint) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement, cdnAuth *CdnAuthConfiguration) error {
	return nil
}

func (i *jsiiProxy_IOriginEndpoint) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IOriginEndpoint) validateMetricEgressBytesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IOriginEndpoint) validateMetricEgressRequestCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IOriginEndpoint) validateMetricEgressResponseTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IOriginEndpoint) validateMetricIngressBytesParameters(options *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IOriginEndpoint) validateMetricIngressRequestCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IOriginEndpoint) validateMetricIngressResponseTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IOriginEndpoint) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

