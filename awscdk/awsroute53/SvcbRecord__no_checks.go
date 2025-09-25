//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SvcbRecord) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SvcbRecord) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SvcbRecord) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSvcbRecord_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSvcbRecord_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSvcbRecord_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSvcbRecordParameters(scope constructs.Construct, id *string, props *SvcbRecordProps) error {
	return nil
}

