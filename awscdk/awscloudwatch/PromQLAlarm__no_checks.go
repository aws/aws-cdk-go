//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PromQLAlarm) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (p *jsiiProxy_PromQLAlarm) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PromQLAlarm) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PromQLAlarm) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePromQLAlarm_FromPromQLAlarmArnParameters(scope constructs.Construct, id *string, alarmArn *string) error {
	return nil
}

func validatePromQLAlarm_FromPromQLAlarmNameParameters(scope constructs.Construct, id *string, alarmName *string) error {
	return nil
}

func validatePromQLAlarm_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePromQLAlarm_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePromQLAlarm_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPromQLAlarmParameters(scope constructs.Construct, id *string, props *PromQLAlarmProps) error {
	return nil
}

