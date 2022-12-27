//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Alarm) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_Alarm) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_Alarm) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAlarm_FromAlarmArnParameters(scope constructs.Construct, id *string, alarmArn *string) error {
	return nil
}

func validateAlarm_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAlarm_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAlarm_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAlarmParameters(scope constructs.Construct, id *string, props *AlarmProps) error {
	return nil
}

