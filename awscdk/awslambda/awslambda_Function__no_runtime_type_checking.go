//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_Function) validateAddAliasParameters(aliasName *string, options *AliasOptions) error {
	return nil
}

func (f *jsiiProxy_Function) validateAddEnvironmentParameters(key *string, value *string, options *EnvironmentOptions) error {
	return nil
}

func (f *jsiiProxy_Function) validateAddEventSourceParameters(source IEventSource) error {
	return nil
}

func (f *jsiiProxy_Function) validateAddEventSourceMappingParameters(id *string, options *EventSourceMappingOptions) error {
	return nil
}

func (f *jsiiProxy_Function) validateAddFunctionUrlParameters(options *FunctionUrlOptions) error {
	return nil
}

func (f *jsiiProxy_Function) validateAddPermissionParameters(id *string, permission *Permission) error {
	return nil
}

func (f *jsiiProxy_Function) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (f *jsiiProxy_Function) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_Function) validateConfigureAsyncInvokeParameters(options *EventInvokeConfigOptions) error {
	return nil
}

func (f *jsiiProxy_Function) validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope constructs.Construct, action *string) error {
	return nil
}

func (f *jsiiProxy_Function) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_Function) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (f *jsiiProxy_Function) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (f *jsiiProxy_Function) validateGrantInvokeUrlParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (f *jsiiProxy_Function) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (f *jsiiProxy_Function) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (f *jsiiProxy_Function) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (f *jsiiProxy_Function) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (f *jsiiProxy_Function) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (f *jsiiProxy_Function) validateWarnInvokeFunctionPermissionsParameters(scope constructs.Construct) error {
	return nil
}

func validateFunction_ClassifyVersionPropertyParameters(propertyName *string, locked *bool) error {
	return nil
}

func validateFunction_FromFunctionArnParameters(scope constructs.Construct, id *string, functionArn *string) error {
	return nil
}

func validateFunction_FromFunctionAttributesParameters(scope constructs.Construct, id *string, attrs *FunctionAttributes) error {
	return nil
}

func validateFunction_FromFunctionNameParameters(scope constructs.Construct, id *string, functionName *string) error {
	return nil
}

func validateFunction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFunction_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFunction_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFunction_MetricAllParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateFunction_MetricAllConcurrentExecutionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateFunction_MetricAllDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateFunction_MetricAllErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateFunction_MetricAllInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateFunction_MetricAllThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateFunction_MetricAllUnreservedConcurrentExecutionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNewFunctionParameters(scope constructs.Construct, id *string, props *FunctionProps) error {
	return nil
}

