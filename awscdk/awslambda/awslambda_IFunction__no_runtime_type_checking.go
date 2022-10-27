//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IFunction) validateAddEventSourceParameters(source IEventSource) error {
	return nil
}

func (i *jsiiProxy_IFunction) validateAddEventSourceMappingParameters(id *string, options *EventSourceMappingOptions) error {
	return nil
}

func (i *jsiiProxy_IFunction) validateAddFunctionUrlParameters(options *FunctionUrlOptions) error {
	return nil
}

func (i *jsiiProxy_IFunction) validateAddPermissionParameters(id *string, permission *Permission) error {
	return nil
}

func (i *jsiiProxy_IFunction) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IFunction) validateConfigureAsyncInvokeParameters(options *EventInvokeConfigOptions) error {
	return nil
}

func (i *jsiiProxy_IFunction) validateGrantInvokeParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IFunction) validateGrantInvokeUrlParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IFunction) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IFunction) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IFunction) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IFunction) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IFunction) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IFunction) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

