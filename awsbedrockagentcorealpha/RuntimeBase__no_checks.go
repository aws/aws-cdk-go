//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RuntimeBase) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateGrantParameters(actions *[]*string, resources *[]*string) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateGrantInvokeRuntimeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateGrantInvokeRuntimeForUserParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateMetricParameters(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateMetricInvocationsAggregatedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateMetricLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateMetricSessionCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateMetricSessionsAggregatedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateMetricSystemErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateMetricTotalErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (r *jsiiProxy_RuntimeBase) validateMetricUserErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateRuntimeBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRuntimeBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRuntimeBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRuntimeBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

