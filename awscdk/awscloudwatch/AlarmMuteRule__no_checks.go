//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlarmMuteRule) validateAddAlarmParameters(alarm interfacesawscloudwatch.IAlarmRef) error {
	return nil
}

func (a *jsiiProxy_AlarmMuteRule) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (a *jsiiProxy_AlarmMuteRule) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AlarmMuteRule) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AlarmMuteRule) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAlarmMuteRule_FromAlarmMuteRuleArnParameters(scope constructs.Construct, id *string, alarmMuteRuleArn *string) error {
	return nil
}

func validateAlarmMuteRule_FromAlarmMuteRuleNameParameters(scope constructs.Construct, id *string, alarmMuteRuleName *string) error {
	return nil
}

func validateAlarmMuteRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAlarmMuteRule_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAlarmMuteRule_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAlarmMuteRuleParameters(scope constructs.Construct, id *string, props *AlarmMuteRuleProps) error {
	return nil
}

