//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CnameRecord) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CnameRecord) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CnameRecord) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCnameRecord_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCnameRecord_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCnameRecord_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCnameRecordParameters(scope constructs.Construct, id *string, props *CnameRecordProps) error {
	return nil
}

