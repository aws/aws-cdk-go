//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RequestAuthorizer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RequestAuthorizer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RequestAuthorizer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateRequestAuthorizer_IsAuthorizerParameters(x interface{}) error {
	return nil
}

func validateRequestAuthorizer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRequestAuthorizer_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRequestAuthorizer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRequestAuthorizerParameters(scope constructs.Construct, id *string, props *RequestAuthorizerProps) error {
	return nil
}

