//go:build no_runtime_type_checking

package awscodecommit

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IRepository) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateGrantPullParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateGrantPullPushParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateNotifyOnParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *RepositoryNotifyOnOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateNotifyOnApprovalRuleOverriddenParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateNotifyOnApprovalStatusChangedParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateNotifyOnBranchOrTagCreatedParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateNotifyOnBranchOrTagDeletedParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateNotifyOnPullRequestCommentParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateNotifyOnPullRequestCreatedParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateNotifyOnPullRequestMergedParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateOnCommentOnCommitParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateOnCommentOnPullRequestParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateOnCommitParameters(id *string, options *OnCommitOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateOnPullRequestStateChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateOnReferenceCreatedParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateOnReferenceDeletedParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateOnReferenceUpdatedParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateOnStateChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateBindAsNotificationRuleSourceParameters(scope constructs.Construct) error {
	return nil
}

