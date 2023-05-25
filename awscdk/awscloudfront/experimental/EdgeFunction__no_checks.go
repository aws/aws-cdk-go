//go:build no_runtime_type_checking

package experimental

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EdgeFunction) validateAddAliasParameters(aliasName *string, options *awslambda.AliasOptions) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateAddEventSourceParameters(source awslambda.IEventSource) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateAddEventSourceMappingParameters(id *string, options *awslambda.EventSourceMappingOptions) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateAddFunctionUrlParameters(options *awslambda.FunctionUrlOptions) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateAddPermissionParameters(id *string, permission *awslambda.Permission) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateConfigureAsyncInvokeParameters(options *awslambda.EventInvokeConfigOptions) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateGrantInvokeParameters(identity awsiam.IGrantable) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateGrantInvokeUrlParameters(identity awsiam.IGrantable) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EdgeFunction) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateEdgeFunction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEdgeFunction_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEdgeFunction_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEdgeFunctionParameters(scope constructs.Construct, id *string, props *EdgeFunctionProps) error {
	return nil
}

