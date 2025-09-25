//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpsRecord) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (h *jsiiProxy_HttpsRecord) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (h *jsiiProxy_HttpsRecord) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateHttpsRecord_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHttpsRecord_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateHttpsRecord_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewHttpsRecordParameters(scope constructs.Construct, id *string, props *HttpsRecordProps) error {
	return nil
}

