//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CaaRecord) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CaaRecord) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CaaRecord) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCaaRecord_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCaaRecord_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCaaRecord_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCaaRecordParameters(scope constructs.Construct, id *string, props *CaaRecordProps) error {
	return nil
}

