//go:build !no_runtime_type_checking

package awsamplify

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (b *jsiiProxy_BasicAuth) validateBindParameters(scope awscdk.Construct, id *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func validateBasicAuth_FromCredentialsParameters(username *string, password awscdk.SecretValue) error {
	if username == nil {
		return fmt.Errorf("parameter username is required, but nil was provided")
	}

	if password == nil {
		return fmt.Errorf("parameter password is required, but nil was provided")
	}

	return nil
}

func validateBasicAuth_FromGeneratedPasswordParameters(username *string) error {
	if username == nil {
		return fmt.Errorf("parameter username is required, but nil was provided")
	}

	return nil
}

func validateNewBasicAuthParameters(props *BasicAuthProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

