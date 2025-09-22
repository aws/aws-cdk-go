//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAlias) validateAddEventSourceParameters(source IEventSource) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateAddEventSourceMappingParameters(id *string, options *EventSourceMappingOptions) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateAddFunctionUrlParameters(options *FunctionUrlOptions) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateAddPermissionParameters(id *string, permission *Permission) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateConfigureAsyncInvokeParameters(options *EventInvokeConfigOptions) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateGrantInvokeParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateGrantInvokeCompositePrincipalParameters(compositePrincipal awsiam.CompositePrincipal) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateGrantInvokeLatestVersionParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateGrantInvokeUrlParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateGrantInvokeVersionParameters(identity awsiam.IGrantable, version IVersion) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

