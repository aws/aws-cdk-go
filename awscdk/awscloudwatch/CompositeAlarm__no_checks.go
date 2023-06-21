//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CompositeAlarm) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CompositeAlarm) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CompositeAlarm) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCompositeAlarm_FromCompositeAlarmArnParameters(scope constructs.Construct, id *string, compositeAlarmArn *string) error {
	return nil
}

func validateCompositeAlarm_FromCompositeAlarmNameParameters(scope constructs.Construct, id *string, compositeAlarmName *string) error {
	return nil
}

func validateCompositeAlarm_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCompositeAlarm_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCompositeAlarm_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCompositeAlarmParameters(scope constructs.Construct, id *string, props *CompositeAlarmProps) error {
	return nil
}

