//go:build no_runtime_type_checking

package awssns

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Subscription) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_Subscription) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Subscription) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSubscription_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSubscription_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSubscription_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSubscriptionParameters(scope constructs.Construct, id *string, props *SubscriptionProps) error {
	return nil
}

