//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

func validateSnapshotCredentials_FromGeneratedPasswordParameters(username *string, options *SnapshotCredentialsFromGeneratedPasswordOptions) error {
	if username == nil {
		return fmt.Errorf("parameter username is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateSnapshotCredentials_FromGeneratedSecretParameters(username *string, options *SnapshotCredentialsFromGeneratedPasswordOptions) error {
	if username == nil {
		return fmt.Errorf("parameter username is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateSnapshotCredentials_FromPasswordParameters(password awscdk.SecretValue) error {
	if password == nil {
		return fmt.Errorf("parameter password is required, but nil was provided")
	}

	return nil
}

func validateSnapshotCredentials_FromSecretParameters(secret awssecretsmanager.ISecret) error {
	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	return nil
}

