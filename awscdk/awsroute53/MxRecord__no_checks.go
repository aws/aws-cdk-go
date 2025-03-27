//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MxRecord) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (m *jsiiProxy_MxRecord) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (m *jsiiProxy_MxRecord) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateMxRecord_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMxRecord_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateMxRecord_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewMxRecordParameters(scope constructs.Construct, id *string, props *MxRecordProps) error {
	return nil
}

