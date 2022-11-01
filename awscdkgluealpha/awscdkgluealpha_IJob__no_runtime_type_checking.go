//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::Glue
package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IJob) validateMetricParameters(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IJob) validateMetricFailureParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IJob) validateMetricSuccessParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IJob) validateMetricTimeoutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IJob) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IJob) validateOnFailureParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IJob) validateOnStateChangeParameters(id *string, jobState JobState, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IJob) validateOnSuccessParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IJob) validateOnTimeoutParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IJob) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

