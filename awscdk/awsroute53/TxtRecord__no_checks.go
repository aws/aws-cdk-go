//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TxtRecord) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TxtRecord) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TxtRecord) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateTxtRecord_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTxtRecord_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTxtRecord_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTxtRecordParameters(scope constructs.Construct, id *string, props *TxtRecordProps) error {
	return nil
}

