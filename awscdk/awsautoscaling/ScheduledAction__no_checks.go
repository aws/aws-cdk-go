//go:build no_runtime_type_checking

package awsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScheduledAction) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_ScheduledAction) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_ScheduledAction) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateScheduledAction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateScheduledAction_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateScheduledAction_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewScheduledActionParameters(scope constructs.Construct, id *string, props *ScheduledActionProps) error {
	return nil
}

