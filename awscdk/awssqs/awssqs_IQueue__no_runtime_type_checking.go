//go:build no_runtime_type_checking

package awssqs

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IQueue) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IQueue) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IQueue) validateGrantConsumeMessagesParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IQueue) validateGrantPurgeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IQueue) validateGrantSendMessagesParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IQueue) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IQueue) validateMetricApproximateAgeOfOldestMessageParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IQueue) validateMetricApproximateNumberOfMessagesDelayedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IQueue) validateMetricApproximateNumberOfMessagesNotVisibleParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IQueue) validateMetricApproximateNumberOfMessagesVisibleParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IQueue) validateMetricNumberOfEmptyReceivesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IQueue) validateMetricNumberOfMessagesDeletedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IQueue) validateMetricNumberOfMessagesReceivedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IQueue) validateMetricNumberOfMessagesSentParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IQueue) validateMetricSentMessageSizeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

