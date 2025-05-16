//go:build no_runtime_type_checking

package awssqs

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_Queue) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (q *jsiiProxy_Queue) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (q *jsiiProxy_Queue) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (q *jsiiProxy_Queue) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (q *jsiiProxy_Queue) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (q *jsiiProxy_Queue) validateGrantConsumeMessagesParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (q *jsiiProxy_Queue) validateGrantPurgeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (q *jsiiProxy_Queue) validateGrantSendMessagesParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (q *jsiiProxy_Queue) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_Queue) validateMetricApproximateAgeOfOldestMessageParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_Queue) validateMetricApproximateNumberOfMessagesDelayedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_Queue) validateMetricApproximateNumberOfMessagesNotVisibleParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_Queue) validateMetricApproximateNumberOfMessagesVisibleParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_Queue) validateMetricNumberOfEmptyReceivesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_Queue) validateMetricNumberOfMessagesDeletedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_Queue) validateMetricNumberOfMessagesReceivedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_Queue) validateMetricNumberOfMessagesSentParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (q *jsiiProxy_Queue) validateMetricSentMessageSizeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateQueue_FromQueueArnParameters(scope constructs.Construct, id *string, queueArn *string) error {
	return nil
}

func validateQueue_FromQueueAttributesParameters(scope constructs.Construct, id *string, attrs *QueueAttributes) error {
	return nil
}

func validateQueue_IsConstructParameters(x interface{}) error {
	return nil
}

func validateQueue_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateQueue_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewQueueParameters(scope constructs.Construct, id *string, props *QueueProps) error {
	return nil
}

