//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PromptBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PromptBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PromptBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_PromptBase) validateGrantGetParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validatePromptBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePromptBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePromptBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPromptBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

