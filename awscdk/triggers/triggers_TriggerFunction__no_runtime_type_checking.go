//go:build no_runtime_type_checking

package triggers

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TriggerFunction) validateAddAliasParameters(aliasName *string, options *awslambda.AliasOptions) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateAddEnvironmentParameters(key *string, value *string, options *awslambda.EnvironmentOptions) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateAddEventSourceParameters(source awslambda.IEventSource) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateAddEventSourceMappingParameters(id *string, options *awslambda.EventSourceMappingOptions) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateAddFunctionUrlParameters(options *awslambda.FunctionUrlOptions) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateAddPermissionParameters(id *string, permission *awslambda.Permission) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateConfigureAsyncInvokeParameters(options *awslambda.EventInvokeConfigOptions) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope constructs.Construct, action *string) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateGrantInvokeUrlParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TriggerFunction) validateWarnInvokeFunctionPermissionsParameters(scope constructs.Construct) error {
	return nil
}

func validateTriggerFunction_ClassifyVersionPropertyParameters(propertyName *string, locked *bool) error {
	return nil
}

func validateTriggerFunction_FromFunctionArnParameters(scope constructs.Construct, id *string, functionArn *string) error {
	return nil
}

func validateTriggerFunction_FromFunctionAttributesParameters(scope constructs.Construct, id *string, attrs *awslambda.FunctionAttributes) error {
	return nil
}

func validateTriggerFunction_FromFunctionNameParameters(scope constructs.Construct, id *string, functionName *string) error {
	return nil
}

func validateTriggerFunction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTriggerFunction_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTriggerFunction_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTriggerFunction_MetricAllParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateTriggerFunction_MetricAllConcurrentExecutionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateTriggerFunction_MetricAllDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateTriggerFunction_MetricAllErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateTriggerFunction_MetricAllInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateTriggerFunction_MetricAllThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateTriggerFunction_MetricAllUnreservedConcurrentExecutionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNewTriggerFunctionParameters(scope constructs.Construct, id *string, props *TriggerFunctionProps) error {
	return nil
}

