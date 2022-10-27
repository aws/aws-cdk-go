//go:build no_runtime_type_checking

package awschatbot

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ISlackChannelConfiguration) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_ISlackChannelConfiguration) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_ISlackChannelConfiguration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_ISlackChannelConfiguration) validateBindAsNotificationRuleTargetParameters(scope constructs.Construct) error {
	return nil
}

