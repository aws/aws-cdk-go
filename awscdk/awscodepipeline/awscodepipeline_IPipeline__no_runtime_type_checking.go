//go:build no_runtime_type_checking

package awscodepipeline

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IPipeline) validateNotifyOnParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *PipelineNotifyOnOptions) error {
	return nil
}

func (i *jsiiProxy_IPipeline) validateNotifyOnAnyActionStateChangeParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (i *jsiiProxy_IPipeline) validateNotifyOnAnyManualApprovalStateChangeParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (i *jsiiProxy_IPipeline) validateNotifyOnAnyStageStateChangeParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (i *jsiiProxy_IPipeline) validateNotifyOnExecutionStateChangeParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (i *jsiiProxy_IPipeline) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IPipeline) validateOnStateChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IPipeline) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IPipeline) validateBindAsNotificationRuleSourceParameters(scope constructs.Construct) error {
	return nil
}

