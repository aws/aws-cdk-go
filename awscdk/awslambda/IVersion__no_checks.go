//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IVersion) validateAddAliasParameters(aliasName *string, options *AliasOptions) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateAddEventSourceParameters(source IEventSource) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateAddEventSourceMappingParameters(id *string, options *EventSourceMappingOptions) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateAddFunctionUrlParameters(options *FunctionUrlOptions) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateAddPermissionParameters(id *string, permission *Permission) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateConfigureAsyncInvokeParameters(options *EventInvokeConfigOptions) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateGrantInvokeParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateGrantInvokeCompositePrincipalParameters(compositePrincipal awsiam.CompositePrincipal) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateGrantInvokeLatestVersionParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateGrantInvokeUrlParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateGrantInvokeVersionParameters(identity awsiam.IGrantable, version IVersion) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IVersion) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

