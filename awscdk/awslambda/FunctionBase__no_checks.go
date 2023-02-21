//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionBase) validateAddEventSourceParameters(source IEventSource) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateAddEventSourceMappingParameters(id *string, options *EventSourceMappingOptions) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateAddFunctionUrlParameters(options *FunctionUrlOptions) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateAddPermissionParameters(id *string, permission *Permission) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateConfigureAsyncInvokeParameters(options *EventInvokeConfigOptions) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope constructs.Construct, action *string) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateGrantInvokeUrlParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (f *jsiiProxy_FunctionBase) validateWarnInvokeFunctionPermissionsParameters(scope constructs.Construct) error {
	return nil
}

func validateFunctionBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFunctionBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFunctionBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFunctionBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

