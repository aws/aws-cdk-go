//go:build no_runtime_type_checking

package awslambdanodejs

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NodejsFunction) validateAddAliasParameters(aliasName *string, options *awslambda.AliasOptions) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateAddEnvironmentParameters(key *string, value *string, options *awslambda.EnvironmentOptions) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateAddEventSourceParameters(source awslambda.IEventSource) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateAddEventSourceMappingParameters(id *string, options *awslambda.EventSourceMappingOptions) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateAddFunctionUrlParameters(options *awslambda.FunctionUrlOptions) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateAddPermissionParameters(id *string, permission *awslambda.Permission) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateConfigureAsyncInvokeParameters(options *awslambda.EventInvokeConfigOptions) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope constructs.Construct, action *string) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateGrantInvokeUrlParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NodejsFunction) validateWarnInvokeFunctionPermissionsParameters(scope constructs.Construct) error {
	return nil
}

func validateNodejsFunction_ClassifyVersionPropertyParameters(propertyName *string, locked *bool) error {
	return nil
}

func validateNodejsFunction_FromFunctionArnParameters(scope constructs.Construct, id *string, functionArn *string) error {
	return nil
}

func validateNodejsFunction_FromFunctionAttributesParameters(scope constructs.Construct, id *string, attrs *awslambda.FunctionAttributes) error {
	return nil
}

func validateNodejsFunction_FromFunctionNameParameters(scope constructs.Construct, id *string, functionName *string) error {
	return nil
}

func validateNodejsFunction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNodejsFunction_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNodejsFunction_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNodejsFunction_MetricAllParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNodejsFunction_MetricAllConcurrentExecutionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNodejsFunction_MetricAllDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNodejsFunction_MetricAllErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNodejsFunction_MetricAllInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNodejsFunction_MetricAllThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNodejsFunction_MetricAllUnreservedConcurrentExecutionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNewNodejsFunctionParameters(scope constructs.Construct, id *string, props *NodejsFunctionProps) error {
	return nil
}

