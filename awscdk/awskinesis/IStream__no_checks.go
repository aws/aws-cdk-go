//go:build no_runtime_type_checking

package awskinesis

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IStream) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IStream) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IStream) validateGrantReadWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IStream) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricGetRecordsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricGetRecordsBytesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricGetRecordsIteratorAgeMillisecondsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricGetRecordsLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricGetRecordsSuccessParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricIncomingBytesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricIncomingRecordsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricPutRecordBytesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricPutRecordLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricPutRecordsBytesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricPutRecordsFailedRecordsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricPutRecordsLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricPutRecordsSuccessParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricPutRecordsSuccessfulRecordsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricPutRecordsThrottledRecordsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricPutRecordsTotalRecordsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricPutRecordSuccessParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricReadProvisionedThroughputExceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IStream) validateMetricWriteProvisionedThroughputExceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

