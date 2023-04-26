//go:build no_runtime_type_checking

package awscognitoidentitypool

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IUserPoolAuthenticationProvider) validateBindParameters(scope constructs.Construct, identityPool IIdentityPool, options *UserPoolAuthenticationProviderBindOptions) error {
	return nil
}

