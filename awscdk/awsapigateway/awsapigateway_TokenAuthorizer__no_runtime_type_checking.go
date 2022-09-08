//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TokenAuthorizer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TokenAuthorizer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TokenAuthorizer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_TokenAuthorizer) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (t *jsiiProxy_TokenAuthorizer) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateTokenAuthorizer_IsAuthorizerParameters(x interface{}) error {
	return nil
}

func validateTokenAuthorizer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTokenAuthorizer_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewTokenAuthorizerParameters(scope constructs.Construct, id *string, props *TokenAuthorizerProps) error {
	return nil
}

