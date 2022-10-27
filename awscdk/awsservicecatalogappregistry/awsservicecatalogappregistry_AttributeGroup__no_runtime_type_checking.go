//go:build no_runtime_type_checking

package awsservicecatalogappregistry

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AttributeGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AttributeGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AttributeGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_AttributeGroup) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_AttributeGroup) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateAttributeGroup_FromAttributeGroupArnParameters(scope constructs.Construct, id *string, attributeGroupArn *string) error {
	return nil
}

func validateAttributeGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAttributeGroup_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewAttributeGroupParameters(scope constructs.Construct, id *string, props *AttributeGroupProps) error {
	return nil
}

