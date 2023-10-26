//go:build no_runtime_type_checking

package awscdklambdagoalpha

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GoFunction) validateAddAliasParameters(aliasName *string, options *awslambda.AliasOptions) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateAddEnvironmentParameters(key *string, value *string, options *awslambda.EnvironmentOptions) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateAddEventSourceParameters(source awslambda.IEventSource) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateAddEventSourceMappingParameters(id *string, options *awslambda.EventSourceMappingOptions) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateAddFunctionUrlParameters(options *awslambda.FunctionUrlOptions) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateAddPermissionParameters(id *string, permission *awslambda.Permission) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateConfigureAsyncInvokeParameters(options *awslambda.EventInvokeConfigOptions) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope constructs.Construct, action *string) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateGrantInvokeCompositePrincipalParameters(compositePrincipal awsiam.CompositePrincipal) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateGrantInvokeUrlParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateInvalidateVersionBasedOnParameters(x *string) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GoFunction) validateWarnInvokeFunctionPermissionsParameters(scope constructs.Construct) error {
	return nil
}

func validateGoFunction_ClassifyVersionPropertyParameters(propertyName *string, locked *bool) error {
	return nil
}

func validateGoFunction_FromFunctionArnParameters(scope constructs.Construct, id *string, functionArn *string) error {
	return nil
}

func validateGoFunction_FromFunctionAttributesParameters(scope constructs.Construct, id *string, attrs *awslambda.FunctionAttributes) error {
	return nil
}

func validateGoFunction_FromFunctionNameParameters(scope constructs.Construct, id *string, functionName *string) error {
	return nil
}

func validateGoFunction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGoFunction_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGoFunction_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGoFunction_MetricAllParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateGoFunction_MetricAllConcurrentExecutionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateGoFunction_MetricAllDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateGoFunction_MetricAllErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateGoFunction_MetricAllInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateGoFunction_MetricAllThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateGoFunction_MetricAllUnreservedConcurrentExecutionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNewGoFunctionParameters(scope constructs.Construct, id *string, props *GoFunctionProps) error {
	return nil
}

