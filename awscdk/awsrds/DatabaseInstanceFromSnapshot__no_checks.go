//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateAddProxyParameters(id *string, options *DatabaseProxyOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateAddRotationMultiUserParameters(id *string, options *RotationMultiUserOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateAddRotationSingleUserParameters(options *RotationSingleUserOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateMetricCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateMetricDatabaseConnectionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateMetricFreeableMemoryParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateMetricFreeStorageSpaceParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateMetricReadIOPSParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateMetricWriteIOPSParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceFromSnapshot) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateDatabaseInstanceFromSnapshot_FromDatabaseInstanceAttributesParameters(scope constructs.Construct, id *string, attrs *DatabaseInstanceAttributes) error {
	return nil
}

func validateDatabaseInstanceFromSnapshot_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseInstanceFromSnapshot_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDatabaseInstanceFromSnapshot_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDatabaseInstanceFromSnapshotParameters(scope constructs.Construct, id *string, props *DatabaseInstanceFromSnapshotProps) error {
	return nil
}

