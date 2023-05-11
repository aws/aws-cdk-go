//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDatabaseCluster) validateAddProxyParameters(id *string, options *DatabaseProxyOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricDatabaseConnectionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricDeadlocksParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricEngineUptimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricFreeableMemoryParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricFreeLocalStorageParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricNetworkReceiveThroughputParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricNetworkThroughputParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricNetworkTransmitThroughputParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricSnapshotStorageUsedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricTotalBackupStorageBilledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricVolumeBytesUsedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricVolumeReadIOPsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateMetricVolumeWriteIOPsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

