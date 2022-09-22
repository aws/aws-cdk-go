//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscognitoidentitypool

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IdentityPoolRoleAttachment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IdentityPoolRoleAttachment) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IdentityPoolRoleAttachment) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (i *jsiiProxy_IdentityPoolRoleAttachment) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (i *jsiiProxy_IdentityPoolRoleAttachment) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateIdentityPoolRoleAttachment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIdentityPoolRoleAttachment_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewIdentityPoolRoleAttachmentParameters(scope constructs.Construct, id *string, props *IdentityPoolRoleAttachmentProps) error {
	return nil
}

