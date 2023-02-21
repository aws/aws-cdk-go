//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IUserPool) validateAddClientParameters(id *string, options *UserPoolClientOptions) error {
	return nil
}

func (i *jsiiProxy_IUserPool) validateAddDomainParameters(id *string, options *UserPoolDomainOptions) error {
	return nil
}

func (i *jsiiProxy_IUserPool) validateAddResourceServerParameters(id *string, options *UserPoolResourceServerOptions) error {
	return nil
}

func (i *jsiiProxy_IUserPool) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IUserPool) validateRegisterIdentityProviderParameters(provider IUserPoolIdentityProvider) error {
	return nil
}

