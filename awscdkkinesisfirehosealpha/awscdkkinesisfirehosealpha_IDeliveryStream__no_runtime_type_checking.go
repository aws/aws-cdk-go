//go:build no_runtime_type_checking

// The CDK Construct Library for AWS::KinesisFirehose
package awscdkkinesisfirehosealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDeliveryStream) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IDeliveryStream) validateGrantPutRecordsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IDeliveryStream) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDeliveryStream) validateMetricBackupToS3BytesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDeliveryStream) validateMetricBackupToS3DataFreshnessParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDeliveryStream) validateMetricBackupToS3RecordsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDeliveryStream) validateMetricIncomingBytesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDeliveryStream) validateMetricIncomingRecordsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IDeliveryStream) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

