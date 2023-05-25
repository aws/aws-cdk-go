//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Alias) validateAddAutoScalingParameters(options *AutoScalingOptions) error {
	return nil
}

func (a *jsiiProxy_Alias) validateAddEventSourceParameters(source IEventSource) error {
	return nil
}

func (a *jsiiProxy_Alias) validateAddEventSourceMappingParameters(id *string, options *EventSourceMappingOptions) error {
	return nil
}

func (a *jsiiProxy_Alias) validateAddFunctionUrlParameters(options *FunctionUrlOptions) error {
	return nil
}

func (a *jsiiProxy_Alias) validateAddPermissionParameters(id *string, permission *Permission) error {
	return nil
}

func (a *jsiiProxy_Alias) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (a *jsiiProxy_Alias) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_Alias) validateConfigureAsyncInvokeParameters(options *EventInvokeConfigOptions) error {
	return nil
}

func (a *jsiiProxy_Alias) validateConsiderWarningOnInvokeFunctionPermissionsParameters(_scope constructs.Construct, _action *string) error {
	return nil
}

func (a *jsiiProxy_Alias) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_Alias) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_Alias) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_Alias) validateGrantInvokeUrlParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_Alias) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Alias) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Alias) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Alias) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Alias) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (a *jsiiProxy_Alias) validateWarnInvokeFunctionPermissionsParameters(scope constructs.Construct) error {
	return nil
}

func validateAlias_FromAliasAttributesParameters(scope constructs.Construct, id *string, attrs *AliasAttributes) error {
	return nil
}

func validateAlias_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAlias_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAlias_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAliasParameters(scope constructs.Construct, id *string, props *AliasProps) error {
	return nil
}

