//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseClusterBase) validateAddProxyParameters(id *string, options *DatabaseProxyOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricDatabaseConnectionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricDeadlocksParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricEngineUptimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricFreeableMemoryParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricFreeLocalStorageParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricNetworkReceiveThroughputParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricNetworkThroughputParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricNetworkTransmitThroughputParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricSnapshotStorageUsedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricTotalBackupStorageBilledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricVolumeBytesUsedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricVolumeReadIOPsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateMetricVolumeWriteIOPsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateDatabaseClusterBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseClusterBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDatabaseClusterBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDatabaseClusterBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

