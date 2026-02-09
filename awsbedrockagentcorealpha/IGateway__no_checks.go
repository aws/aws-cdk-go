//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IGateway) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IGateway) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IGateway) validateGrantManageParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IGateway) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IGateway) validateMetricParameters(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGateway) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGateway) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGateway) validateMetricLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGateway) validateMetricSystemErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGateway) validateMetricTargetExecutionTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGateway) validateMetricTargetTypeParameters(targetType *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGateway) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGateway) validateMetricUserErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGateway) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

