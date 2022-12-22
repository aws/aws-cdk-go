//go:build no_runtime_type_checking

package awsamplify

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BasicAuth) validateBindParameters(scope awscdk.Construct, id *string) error {
	return nil
}

func validateBasicAuth_FromCredentialsParameters(username *string, password awscdk.SecretValue) error {
	return nil
}

func validateBasicAuth_FromGeneratedPasswordParameters(username *string) error {
	return nil
}

func validateNewBasicAuthParameters(props *BasicAuthProps) error {
	return nil
}

