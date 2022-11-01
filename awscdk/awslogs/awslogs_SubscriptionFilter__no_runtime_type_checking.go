//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SubscriptionFilter) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SubscriptionFilter) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SubscriptionFilter) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_SubscriptionFilter) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_SubscriptionFilter) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateSubscriptionFilter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSubscriptionFilter_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewSubscriptionFilterParameters(scope constructs.Construct, id *string, props *SubscriptionFilterProps) error {
	return nil
}

