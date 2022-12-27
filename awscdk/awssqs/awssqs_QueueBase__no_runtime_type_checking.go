//go:build no_runtime_type_checking

package awssqs

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueueBase) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateGrantConsumeMessagesParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateGrantPurgeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateGrantSendMessagesParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateMetricApproximateAgeOfOldestMessageParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateMetricApproximateNumberOfMessagesDelayedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateMetricApproximateNumberOfMessagesNotVisibleParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateMetricApproximateNumberOfMessagesVisibleParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateMetricNumberOfEmptyReceivesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateMetricNumberOfMessagesDeletedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateMetricNumberOfMessagesReceivedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateMetricNumberOfMessagesSentParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_QueueBase) validateMetricSentMessageSizeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateQueueBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateQueueBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateQueueBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewQueueBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

