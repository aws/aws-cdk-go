//go:build no_runtime_type_checking

package awsservicediscovery

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CnameInstance) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CnameInstance) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CnameInstance) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCnameInstance_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCnameInstance_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCnameInstance_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCnameInstanceParameters(scope constructs.Construct, id *string, props *CnameInstanceProps) error {
	return nil
}

