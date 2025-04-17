//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SrvRecord) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SrvRecord) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SrvRecord) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSrvRecord_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSrvRecord_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSrvRecord_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSrvRecordParameters(scope constructs.Construct, id *string, props *SrvRecordProps) error {
	return nil
}

