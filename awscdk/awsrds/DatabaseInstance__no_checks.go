//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseInstance) validateAddProxyParameters(id *string, options *DatabaseProxyOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateAddRotationMultiUserParameters(id *string, options *RotationMultiUserOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateAddRotationSingleUserParameters(options *RotationSingleUserOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateMetricCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateMetricDatabaseConnectionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateMetricFreeableMemoryParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateMetricFreeStorageSpaceParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateMetricReadIOPSParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateMetricWriteIOPSParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateDatabaseInstance_FromDatabaseInstanceAttributesParameters(scope constructs.Construct, id *string, attrs *DatabaseInstanceAttributes) error {
	return nil
}

func validateDatabaseInstance_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseInstance_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDatabaseInstance_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDatabaseInstanceParameters(scope constructs.Construct, id *string, props *DatabaseInstanceProps) error {
	return nil
}

