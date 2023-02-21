//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NsRecord) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NsRecord) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NsRecord) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateNsRecord_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNsRecord_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNsRecord_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNsRecordParameters(scope constructs.Construct, id *string, props *NsRecordProps) error {
	return nil
}

