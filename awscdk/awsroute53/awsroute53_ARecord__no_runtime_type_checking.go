//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ARecord) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_ARecord) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_ARecord) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateARecord_IsConstructParameters(x interface{}) error {
	return nil
}

func validateARecord_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateARecord_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewARecordParameters(scope constructs.Construct, id *string, props *ARecordProps) error {
	return nil
}

