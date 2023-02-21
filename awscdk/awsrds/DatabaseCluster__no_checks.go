//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseCluster) validateAddProxyParameters(id *string, options *DatabaseProxyOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateAddRotationMultiUserParameters(id *string, options *RotationMultiUserOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateAddRotationSingleUserParameters(options *RotationSingleUserOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricDatabaseConnectionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricDeadlocksParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricEngineUptimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricFreeableMemoryParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricFreeLocalStorageParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricNetworkReceiveThroughputParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricNetworkThroughputParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricNetworkTransmitThroughputParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricSnapshotStorageUsedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricTotalBackupStorageBilledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricVolumeBytesUsedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricVolumeReadIOPsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateMetricVolumeWriteIOPsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateDatabaseCluster_FromDatabaseClusterAttributesParameters(scope constructs.Construct, id *string, attrs *DatabaseClusterAttributes) error {
	return nil
}

func validateDatabaseCluster_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseCluster_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDatabaseCluster_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDatabaseClusterParameters(scope constructs.Construct, id *string, props *DatabaseClusterProps) error {
	return nil
}

