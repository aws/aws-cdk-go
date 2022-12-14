//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DockerImageFunction) validateAddAliasParameters(aliasName *string, options *AliasOptions) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateAddEnvironmentParameters(key *string, value *string, options *EnvironmentOptions) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateAddEventSourceParameters(source IEventSource) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateAddEventSourceMappingParameters(id *string, options *EventSourceMappingOptions) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateAddFunctionUrlParameters(options *FunctionUrlOptions) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateAddPermissionParameters(id *string, permission *Permission) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateConfigureAsyncInvokeParameters(options *EventInvokeConfigOptions) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope constructs.Construct, action *string) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateGrantInvokeUrlParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DockerImageFunction) validateWarnInvokeFunctionPermissionsParameters(scope constructs.Construct) error {
	return nil
}

func validateDockerImageFunction_ClassifyVersionPropertyParameters(propertyName *string, locked *bool) error {
	return nil
}

func validateDockerImageFunction_FromFunctionArnParameters(scope constructs.Construct, id *string, functionArn *string) error {
	return nil
}

func validateDockerImageFunction_FromFunctionAttributesParameters(scope constructs.Construct, id *string, attrs *FunctionAttributes) error {
	return nil
}

func validateDockerImageFunction_FromFunctionNameParameters(scope constructs.Construct, id *string, functionName *string) error {
	return nil
}

func validateDockerImageFunction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDockerImageFunction_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDockerImageFunction_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDockerImageFunction_MetricAllParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateDockerImageFunction_MetricAllConcurrentExecutionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateDockerImageFunction_MetricAllDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateDockerImageFunction_MetricAllErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateDockerImageFunction_MetricAllInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateDockerImageFunction_MetricAllThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateDockerImageFunction_MetricAllUnreservedConcurrentExecutionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNewDockerImageFunctionParameters(scope constructs.Construct, id *string, props *DockerImageFunctionProps) error {
	return nil
}

