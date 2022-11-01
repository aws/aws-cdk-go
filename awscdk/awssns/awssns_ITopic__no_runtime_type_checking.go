//go:build no_runtime_type_checking

package awssns

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITopic) validateAddSubscriptionParameters(subscription ITopicSubscription) error {
	return nil
}

func (i *jsiiProxy_ITopic) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_ITopic) validateGrantPublishParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ITopic) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITopic) validateMetricNumberOfMessagesPublishedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITopic) validateMetricNumberOfNotificationsDeliveredParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITopic) validateMetricNumberOfNotificationsFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITopic) validateMetricNumberOfNotificationsFilteredOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITopic) validateMetricNumberOfNotificationsFilteredOutInvalidAttributesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITopic) validateMetricNumberOfNotificationsFilteredOutNoMessageAttributesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITopic) validateMetricPublishSizeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITopic) validateMetricSMSMonthToDateSpentUSDParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITopic) validateMetricSMSSuccessRateParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ITopic) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_ITopic) validateBindAsNotificationRuleTargetParameters(scope constructs.Construct) error {
	return nil
}

