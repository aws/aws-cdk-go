//go:build no_runtime_type_checking

package awssns

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TopicBase) validateAddSubscriptionParameters(topicSubscription ITopicSubscription) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateBindAsNotificationRuleTargetParameters(_scope constructs.Construct) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateGrantPublishParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateMetricNumberOfMessagesPublishedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateMetricNumberOfNotificationsDeliveredParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateMetricNumberOfNotificationsFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateMetricNumberOfNotificationsFilteredOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateMetricNumberOfNotificationsFilteredOutInvalidAttributesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateMetricNumberOfNotificationsFilteredOutNoMessageAttributesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateMetricPublishSizeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateMetricSMSMonthToDateSpentUSDParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TopicBase) validateMetricSMSSuccessRateParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateTopicBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTopicBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTopicBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTopicBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

