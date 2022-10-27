//go:build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OpenIdConnectProvider) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (o *jsiiProxy_OpenIdConnectProvider) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (o *jsiiProxy_OpenIdConnectProvider) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (o *jsiiProxy_OpenIdConnectProvider) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (o *jsiiProxy_OpenIdConnectProvider) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateOpenIdConnectProvider_FromOpenIdConnectProviderArnParameters(scope constructs.Construct, id *string, openIdConnectProviderArn *string) error {
	return nil
}

func validateOpenIdConnectProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateOpenIdConnectProvider_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewOpenIdConnectProviderParameters(scope constructs.Construct, id *string, props *OpenIdConnectProviderProps) error {
	return nil
}

