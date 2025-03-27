//go:build !no_runtime_type_checking

package awscdkcognitoidentitypoolalpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
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

func validateIdentityPoolProviderUrl_UserPoolParameters(userPool awscognito.IUserPool, userPoolClient awscognito.IUserPoolClient) error {
	if userPool == nil {
		return fmt.Errorf("parameter userPool is required, but nil was provided")
	}

	if userPoolClient == nil {
		return fmt.Errorf("parameter userPoolClient is required, but nil was provided")
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

