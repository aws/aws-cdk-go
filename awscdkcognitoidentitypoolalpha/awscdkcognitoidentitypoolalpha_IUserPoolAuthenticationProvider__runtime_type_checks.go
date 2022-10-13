//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_IUserPoolAuthenticationProvider) validateBindParameters(scope constructs.Construct, identityPool IIdentityPool, options *UserPoolAuthenticationProviderBindOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if identityPool == nil {
		return fmt.Errorf("parameter identityPool is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

