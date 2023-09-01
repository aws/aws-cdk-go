//go:build no_runtime_type_checking

package awscdkscheduleralpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Schedule) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_Schedule) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Schedule) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSchedule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSchedule_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSchedule_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewScheduleParameters(scope constructs.Construct, id *string, props *ScheduleProps) error {
	return nil
}

