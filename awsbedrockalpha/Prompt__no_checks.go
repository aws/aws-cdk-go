//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Prompt) validateAddVariantParameters(variant IPromptVariant) error {
	return nil
}

func (p *jsiiProxy_Prompt) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_Prompt) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_Prompt) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_Prompt) validateGrantGetParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validatePrompt_FromPromptAttributesParameters(scope constructs.Construct, id *string, attrs *PromptAttributes) error {
	return nil
}

func validatePrompt_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePrompt_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePrompt_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPromptParameters(scope constructs.Construct, id *string, props *PromptProps) error {
	return nil
}

