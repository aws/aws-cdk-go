//go:build no_runtime_type_checking

package awsservicecatalogappregistry

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Application) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_Application) validateAssociateAttributeGroupParameters(attributeGroup IAttributeGroup) error {
	return nil
}

func (a *jsiiProxy_Application) validateAssociateStackParameters(stack awscdk.Stack) error {
	return nil
}

func (a *jsiiProxy_Application) validateGenerateUniqueHashParameters(resourceAddress *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_Application) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_Application) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateApplication_FromApplicationArnParameters(scope constructs.Construct, id *string, applicationArn *string) error {
	return nil
}

func validateApplication_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApplication_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewApplicationParameters(scope constructs.Construct, id *string, props *ApplicationProps) error {
	return nil
}

