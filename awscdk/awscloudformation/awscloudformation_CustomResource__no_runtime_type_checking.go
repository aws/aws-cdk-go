//go:build no_runtime_type_checking

package awscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomResource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetAttStringParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CustomResource) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateCustomResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCustomResource_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewCustomResourceParameters(scope awscdk.Construct, id *string, props *CustomResourceProps) error {
	return nil
}

