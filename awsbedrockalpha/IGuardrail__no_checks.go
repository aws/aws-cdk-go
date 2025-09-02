//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IGuardrail) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IGuardrail) validateGrantApplyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IGuardrail) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGuardrail) validateMetricInvocationClientErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGuardrail) validateMetricInvocationLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGuardrail) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGuardrail) validateMetricInvocationServerErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGuardrail) validateMetricInvocationsIntervenedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGuardrail) validateMetricInvocationThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IGuardrail) validateMetricTextUnitCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

