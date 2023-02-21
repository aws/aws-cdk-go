//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomResource) validateApplyRemovalPolicyParameters(policy RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetAttStringParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCustomResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCustomResource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCustomResource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCustomResourceParameters(scope constructs.Construct, id *string, props *CustomResourceProps) error {
	return nil
}

