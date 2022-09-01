//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha

import (
	"fmt"
)

func validateIdentityPoolProviderUrl_CustomParameters(url *string) error {
	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	return nil
}

func validateIdentityPoolProviderUrl_OpenIdParameters(url *string) error {
	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	return nil
}

func validateIdentityPoolProviderUrl_SamlParameters(url *string) error {
	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	return nil
}

func validateIdentityPoolProviderUrl_UserPoolParameters(url *string) error {
	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	return nil
}

func validateNewIdentityPoolProviderUrlParameters(type_ IdentityPoolProviderType, value *string) error {
	if type_ == "" {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

