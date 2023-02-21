//go:build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EmailIdentity) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EmailIdentity) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EmailIdentity) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateEmailIdentity_FromEmailIdentityNameParameters(scope constructs.Construct, id *string, emailIdentityName *string) error {
	return nil
}

func validateEmailIdentity_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEmailIdentity_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEmailIdentity_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEmailIdentityParameters(scope constructs.Construct, id *string, props *EmailIdentityProps) error {
	return nil
}

