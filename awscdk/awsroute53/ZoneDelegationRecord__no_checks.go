//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (z *jsiiProxy_ZoneDelegationRecord) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (z *jsiiProxy_ZoneDelegationRecord) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (z *jsiiProxy_ZoneDelegationRecord) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateZoneDelegationRecord_IsConstructParameters(x interface{}) error {
	return nil
}

func validateZoneDelegationRecord_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateZoneDelegationRecord_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewZoneDelegationRecordParameters(scope constructs.Construct, id *string, props *ZoneDelegationRecordProps) error {
	return nil
}

