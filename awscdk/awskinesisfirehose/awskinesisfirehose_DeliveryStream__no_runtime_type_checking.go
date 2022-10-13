//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awskinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeliveryStream) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DeliveryStream) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DeliveryStream) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DeliveryStream) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DeliveryStream) validateGrantPutRecordsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DeliveryStream) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DeliveryStream) validateMetricBackupToS3BytesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DeliveryStream) validateMetricBackupToS3DataFreshnessParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DeliveryStream) validateMetricBackupToS3RecordsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DeliveryStream) validateMetricIncomingBytesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DeliveryStream) validateMetricIncomingRecordsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DeliveryStream) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (d *jsiiProxy_DeliveryStream) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateDeliveryStream_FromDeliveryStreamArnParameters(scope constructs.Construct, id *string, deliveryStreamArn *string) error {
	return nil
}

func validateDeliveryStream_FromDeliveryStreamAttributesParameters(scope constructs.Construct, id *string, attrs *DeliveryStreamAttributes) error {
	return nil
}

func validateDeliveryStream_FromDeliveryStreamNameParameters(scope constructs.Construct, id *string, deliveryStreamName *string) error {
	return nil
}

func validateDeliveryStream_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDeliveryStream_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewDeliveryStreamParameters(scope constructs.Construct, id *string, props *DeliveryStreamProps) error {
	return nil
}

