//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseInstanceReadReplica) validateAddProxyParameters(id *string, options *DatabaseProxyOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceReadReplica) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceReadReplica) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceReadReplica) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceReadReplica) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceReadReplica) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceReadReplica) validateMetricCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceReadReplica) validateMetricDatabaseConnectionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceReadReplica) validateMetricFreeableMemoryParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceReadReplica) validateMetricFreeStorageSpaceParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceReadReplica) validateMetricReadIOPSParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceReadReplica) validateMetricWriteIOPSParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceReadReplica) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateDatabaseInstanceReadReplica_FromDatabaseInstanceAttributesParameters(scope constructs.Construct, id *string, attrs *DatabaseInstanceAttributes) error {
	return nil
}

func validateDatabaseInstanceReadReplica_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseInstanceReadReplica_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDatabaseInstanceReadReplica_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDatabaseInstanceReadReplicaParameters(scope constructs.Construct, id *string, props *DatabaseInstanceReadReplicaProps) error {
	return nil
}

