//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AnomalyDetectionAlarm) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AnomalyDetectionAlarm) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AnomalyDetectionAlarm) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAnomalyDetectionAlarm_FromAlarmArnParameters(scope constructs.Construct, id *string, alarmArn *string) error {
	return nil
}

func validateAnomalyDetectionAlarm_FromAlarmNameParameters(scope constructs.Construct, id *string, alarmName *string) error {
	return nil
}

func validateAnomalyDetectionAlarm_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAnomalyDetectionAlarm_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAnomalyDetectionAlarm_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAnomalyDetectionAlarmParameters(scope constructs.Construct, id *string, props *AnomalyDetectionAlarmProps) error {
	return nil
}

