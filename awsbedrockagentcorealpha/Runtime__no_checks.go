//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Runtime) validateAddEndpointParameters(endpointName *string, options *AddEndpointOptions) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateGrantParameters(actions *[]*string, resources *[]*string) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateGrantInvokeRuntimeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateGrantInvokeRuntimeForUserParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateMetricParameters(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateMetricInvocationsAggregatedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateMetricLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateMetricSessionCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateMetricSessionsAggregatedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateMetricSystemErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateMetricTotalErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_Runtime) validateMetricUserErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateRuntime_FromAgentRuntimeAttributesParameters(scope constructs.Construct, id *string, attrs *AgentRuntimeAttributes) error {
	return nil
}

func validateRuntime_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRuntime_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRuntime_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRuntimeParameters(scope constructs.Construct, id *string, props *RuntimeProps) error {
	return nil
}

