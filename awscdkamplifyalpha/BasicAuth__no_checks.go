//go:build no_runtime_type_checking

// The CDK Construct Library for AWS::Amplify
package awscdkamplifyalpha

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BasicAuth) validateBindParameters(scope constructs.Construct, id *string) error {
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

