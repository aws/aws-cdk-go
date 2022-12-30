//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

func validateCredentials_FromGeneratedSecretParameters(username *string, options *CredentialsBaseOptions) error {
	if username == nil {
		return fmt.Errorf("parameter username is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateCredentials_FromPasswordParameters(username *string, password awscdk.SecretValue) error {
	if username == nil {
		return fmt.Errorf("parameter username is required, but nil was provided")
	}

	if password == nil {
		return fmt.Errorf("parameter password is required, but nil was provided")
	}

	return nil
}

func validateCredentials_FromSecretParameters(secret awssecretsmanager.ISecret) error {
	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	return nil
}

func validateCredentials_FromUsernameParameters(username *string, options *CredentialsFromUsernameOptions) error {
	if username == nil {
		return fmt.Errorf("parameter username is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

