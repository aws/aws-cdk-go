//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IApiKeyCredentialProvider) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IApiKeyCredentialProvider) validateGrantAdminParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IApiKeyCredentialProvider) validateGrantFullAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IApiKeyCredentialProvider) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IApiKeyCredentialProvider) validateGrantUseParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IApiKeyCredentialProvider) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

