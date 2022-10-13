//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserPoolAuthenticationProvider) validateBindParameters(scope constructs.Construct, identityPool IIdentityPool, _options *UserPoolAuthenticationProviderBindOptions) error {
	return nil
}

func validateNewUserPoolAuthenticationProviderParameters(props *UserPoolAuthenticationProviderProps) error {
	return nil
}

