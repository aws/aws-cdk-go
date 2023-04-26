//go:build no_runtime_type_checking

package awssns

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Topic) validateAddSubscriptionParameters(topicSubscription ITopicSubscription) error {
	return nil
}

func (t *jsiiProxy_Topic) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (t *jsiiProxy_Topic) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_Topic) validateBindAsNotificationRuleTargetParameters(_scope constructs.Construct) error {
	return nil
}

func (t *jsiiProxy_Topic) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_Topic) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_Topic) validateGrantPublishParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_Topic) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Topic) validateMetricNumberOfMessagesPublishedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Topic) validateMetricNumberOfNotificationsDeliveredParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Topic) validateMetricNumberOfNotificationsFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Topic) validateMetricNumberOfNotificationsFilteredOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Topic) validateMetricNumberOfNotificationsFilteredOutInvalidAttributesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Topic) validateMetricNumberOfNotificationsFilteredOutNoMessageAttributesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Topic) validateMetricPublishSizeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Topic) validateMetricSMSMonthToDateSpentUSDParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Topic) validateMetricSMSSuccessRateParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateTopic_FromTopicArnParameters(scope constructs.Construct, id *string, topicArn *string) error {
	return nil
}

func validateTopic_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTopic_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTopic_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTopicParameters(scope constructs.Construct, id *string, props *TopicProps) error {
	return nil
}

