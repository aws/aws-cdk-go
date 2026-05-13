//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IOAuth2CredentialProvider) validateBindForGatewayOAuthTargetParameters(scopes *[]*string) error {
	return nil
}

func (i *jsiiProxy_IOAuth2CredentialProvider) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IOAuth2CredentialProvider) validateGrantAdminParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IOAuth2CredentialProvider) validateGrantFullAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IOAuth2CredentialProvider) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IOAuth2CredentialProvider) validateGrantUseParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IOAuth2CredentialProvider) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

