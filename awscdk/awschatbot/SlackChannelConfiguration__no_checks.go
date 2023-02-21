//go:build no_runtime_type_checking

package awschatbot

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SlackChannelConfiguration) validateAddNotificationTopicParameters(notificationTopic awssns.ITopic) error {
	return nil
}

func (s *jsiiProxy_SlackChannelConfiguration) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SlackChannelConfiguration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SlackChannelConfiguration) validateBindAsNotificationRuleTargetParameters(_scope constructs.Construct) error {
	return nil
}

func (s *jsiiProxy_SlackChannelConfiguration) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SlackChannelConfiguration) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_SlackChannelConfiguration) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateSlackChannelConfiguration_FromSlackChannelConfigurationArnParameters(scope constructs.Construct, id *string, slackChannelConfigurationArn *string) error {
	return nil
}

func validateSlackChannelConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSlackChannelConfiguration_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSlackChannelConfiguration_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSlackChannelConfiguration_MetricAllParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNewSlackChannelConfigurationParameters(scope constructs.Construct, id *string, props *SlackChannelConfigurationProps) error {
	return nil
}

