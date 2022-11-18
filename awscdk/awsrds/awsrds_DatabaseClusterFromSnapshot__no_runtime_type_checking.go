//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateAddProxyParameters(id *string, options *DatabaseProxyOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateAddRotationMultiUserParameters(id *string, options *RotationMultiUserOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateAddRotationSingleUserParameters(options *RotationSingleUserOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricDatabaseConnectionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricDeadlocksParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricEngineUptimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricFreeableMemoryParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricFreeLocalStorageParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricNetworkReceiveThroughputParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricNetworkThroughputParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricNetworkTransmitThroughputParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricSnapshotStorageUsedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricTotalBackupStorageBilledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricVolumeBytesUsedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricVolumeReadIOPsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterFromSnapshot) validateMetricVolumeWriteIOPsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateDatabaseClusterFromSnapshot_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseClusterFromSnapshot_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDatabaseClusterFromSnapshot_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDatabaseClusterFromSnapshotParameters(scope constructs.Construct, id *string, props *DatabaseClusterFromSnapshotProps) error {
	return nil
}

