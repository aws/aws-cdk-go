//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlarmBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AlarmBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AlarmBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAlarmBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAlarmBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAlarmBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAlarmBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

