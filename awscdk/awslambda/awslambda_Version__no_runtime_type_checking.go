//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_Version) validateAddAliasParameters(aliasName *string, options *AliasOptions) error {
	return nil
}

func (v *jsiiProxy_Version) validateAddEventSourceParameters(source IEventSource) error {
	return nil
}

func (v *jsiiProxy_Version) validateAddEventSourceMappingParameters(id *string, options *EventSourceMappingOptions) error {
	return nil
}

func (v *jsiiProxy_Version) validateAddFunctionUrlParameters(options *FunctionUrlOptions) error {
	return nil
}

func (v *jsiiProxy_Version) validateAddPermissionParameters(id *string, permission *Permission) error {
	return nil
}

func (v *jsiiProxy_Version) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (v *jsiiProxy_Version) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_Version) validateConfigureAsyncInvokeParameters(options *EventInvokeConfigOptions) error {
	return nil
}

func (v *jsiiProxy_Version) validateConsiderWarningOnInvokeFunctionPermissionsParameters(_scope constructs.Construct, _action *string) error {
	return nil
}

func (v *jsiiProxy_Version) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_Version) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (v *jsiiProxy_Version) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (v *jsiiProxy_Version) validateGrantInvokeUrlParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (v *jsiiProxy_Version) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (v *jsiiProxy_Version) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (v *jsiiProxy_Version) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (v *jsiiProxy_Version) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (v *jsiiProxy_Version) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (v *jsiiProxy_Version) validateWarnInvokeFunctionPermissionsParameters(scope constructs.Construct) error {
	return nil
}

func validateVersion_FromVersionArnParameters(scope constructs.Construct, id *string, versionArn *string) error {
	return nil
}

func validateVersion_FromVersionAttributesParameters(scope constructs.Construct, id *string, attrs *VersionAttributes) error {
	return nil
}

func validateVersion_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVersion_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVersion_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVersionParameters(scope constructs.Construct, id *string, props *VersionProps) error {
	return nil
}

