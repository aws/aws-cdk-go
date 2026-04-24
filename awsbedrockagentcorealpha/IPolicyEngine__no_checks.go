//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IPolicyEngine) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IPolicyEngine) validateGrantEvaluateParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IPolicyEngine) validateGrantEvaluateForGatewayParameters(grantee awsiam.IGrantable, gateway IGateway) error {
	return nil
}

func (i *jsiiProxy_IPolicyEngine) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IPolicyEngine) validateMetricParameters(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IPolicyEngine) validateMetricAuthorizationLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IPolicyEngine) validateMetricDeniedRequestsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IPolicyEngine) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IPolicyEngine) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

