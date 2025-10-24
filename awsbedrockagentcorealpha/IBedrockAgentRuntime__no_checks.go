//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IBedrockAgentRuntime) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateGrantParameters(actions *[]*string, resources *[]*string) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateGrantInvokeRuntimeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateGrantInvokeRuntimeForUserParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateMetricParameters(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateMetricInvocationsAggregatedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateMetricLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateMetricSessionCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateMetricSessionsAggregatedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateMetricSystemErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateMetricTotalErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateMetricUserErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IBedrockAgentRuntime) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

