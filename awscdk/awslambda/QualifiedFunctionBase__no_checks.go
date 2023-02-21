//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QualifiedFunctionBase) validateAddEventSourceParameters(source IEventSource) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateAddEventSourceMappingParameters(id *string, options *EventSourceMappingOptions) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateAddFunctionUrlParameters(options *FunctionUrlOptions) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateAddPermissionParameters(id *string, permission *Permission) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateConfigureAsyncInvokeParameters(options *EventInvokeConfigOptions) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateConsiderWarningOnInvokeFunctionPermissionsParameters(_scope constructs.Construct, _action *string) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateGrantInvokeUrlParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_QualifiedFunctionBase) validateWarnInvokeFunctionPermissionsParameters(scope constructs.Construct) error {
	return nil
}

func validateQualifiedFunctionBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateQualifiedFunctionBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateQualifiedFunctionBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewQualifiedFunctionBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

