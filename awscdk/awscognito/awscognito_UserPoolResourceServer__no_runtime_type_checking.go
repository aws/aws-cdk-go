//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserPoolResourceServer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UserPoolResourceServer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UserPoolResourceServer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (u *jsiiProxy_UserPoolResourceServer) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (u *jsiiProxy_UserPoolResourceServer) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateUserPoolResourceServer_FromUserPoolResourceServerIdParameters(scope constructs.Construct, id *string, userPoolResourceServerId *string) error {
	return nil
}

func validateUserPoolResourceServer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUserPoolResourceServer_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewUserPoolResourceServerParameters(scope constructs.Construct, id *string, props *UserPoolResourceServerProps) error {
	return nil
}

