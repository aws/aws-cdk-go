//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDatabaseInstance) validateAddProxyParameters(id *string, options *DatabaseProxyOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseInstance) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IDatabaseInstance) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseInstance) validateMetricCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseInstance) validateMetricDatabaseConnectionsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseInstance) validateMetricFreeableMemoryParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseInstance) validateMetricFreeStorageSpaceParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseInstance) validateMetricReadIOPSParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseInstance) validateMetricWriteIOPSParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseInstance) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IDatabaseInstance) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

