//go:build no_runtime_type_checking

package awscodestarnotifications

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NotificationRule) validateAddTargetParameters(target INotificationRuleTarget) error {
	return nil
}

func (n *jsiiProxy_NotificationRule) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NotificationRule) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NotificationRule) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateNotificationRule_FromNotificationRuleArnParameters(scope constructs.Construct, id *string, notificationRuleArn *string) error {
	return nil
}

func validateNotificationRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNotificationRule_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNotificationRule_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNotificationRuleParameters(scope constructs.Construct, id *string, props *NotificationRuleProps) error {
	return nil
}

