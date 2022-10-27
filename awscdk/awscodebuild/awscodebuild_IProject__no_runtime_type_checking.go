//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IProject) validateAddToRolePolicyParameters(policyStatement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IProject) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IProject) validateMetricBuildsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IProject) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IProject) validateMetricFailedBuildsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IProject) validateMetricSucceededBuildsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IProject) validateNotifyOnParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *ProjectNotifyOnOptions) error {
	return nil
}

func (i *jsiiProxy_IProject) validateNotifyOnBuildFailedParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (i *jsiiProxy_IProject) validateNotifyOnBuildSucceededParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (i *jsiiProxy_IProject) validateOnBuildFailedParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IProject) validateOnBuildStartedParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IProject) validateOnBuildSucceededParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IProject) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IProject) validateOnPhaseChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IProject) validateOnStateChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IProject) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IProject) validateBindAsNotificationRuleSourceParameters(scope constructs.Construct) error {
	return nil
}

