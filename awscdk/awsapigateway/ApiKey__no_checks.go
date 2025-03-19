//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiKey) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_ApiKey) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_ApiKey) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_ApiKey) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_ApiKey) validateGrantReadWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (a *jsiiProxy_ApiKey) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateApiKey_FromApiKeyIdParameters(scope constructs.Construct, id *string, apiKeyId *string) error {
	return nil
}

func validateApiKey_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApiKey_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateApiKey_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewApiKeyParameters(scope constructs.Construct, id *string, props *ApiKeyProps) error {
	return nil
}

