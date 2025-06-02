//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITable) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ITable) validateGrantFullAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ITable) validateGrantReadDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ITable) validateGrantReadWriteDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ITable) validateGrantStreamParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ITable) validateGrantStreamReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ITable) validateGrantTableListStreamsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ITable) validateGrantWriteDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ITable) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITable) validateMetricConditionalCheckFailedRequestsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITable) validateMetricConsumedReadCapacityUnitsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITable) validateMetricConsumedWriteCapacityUnitsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITable) validateMetricSuccessfulRequestLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITable) validateMetricSystemErrorsForOperationsParameters(props *SystemErrorsForOperationsMetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITable) validateMetricThrottledRequestsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITable) validateMetricThrottledRequestsForOperationsParameters(props *OperationsMetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITable) validateMetricUserErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

