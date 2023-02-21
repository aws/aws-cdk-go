//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseInstanceBase) validateAddProxyParameters(id *string, options *DatabaseProxyOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateMetricCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateMetricDatabaseConnectionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateMetricFreeableMemoryParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateMetricFreeStorageSpaceParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateMetricReadIOPSParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateMetricWriteIOPSParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateDatabaseInstanceBase_FromDatabaseInstanceAttributesParameters(scope constructs.Construct, id *string, attrs *DatabaseInstanceAttributes) error {
	return nil
}

func validateDatabaseInstanceBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseInstanceBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDatabaseInstanceBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDatabaseInstanceBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

