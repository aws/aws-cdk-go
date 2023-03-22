//go:build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventBusPolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EventBusPolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EventBusPolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateEventBusPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEventBusPolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEventBusPolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEventBusPolicyParameters(scope constructs.Construct, id *string, props *EventBusPolicyProps) error {
	return nil
}

