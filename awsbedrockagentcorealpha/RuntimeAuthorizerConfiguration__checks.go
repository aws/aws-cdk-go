//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"
)

func validateRuntimeAuthorizerConfiguration_UsingCognitoParameters(userPoolId *string, clientId *string) error {
	if userPoolId == nil {
		return fmt.Errorf("parameter userPoolId is required, but nil was provided")
	}

	if clientId == nil {
		return fmt.Errorf("parameter clientId is required, but nil was provided")
	}

	return nil
}

func validateRuntimeAuthorizerConfiguration_UsingJWTParameters(discoveryUrl *string) error {
	if discoveryUrl == nil {
		return fmt.Errorf("parameter discoveryUrl is required, but nil was provided")
	}

	return nil
}

func validateRuntimeAuthorizerConfiguration_UsingOAuthParameters(discoveryUrl *string, clientId *string) error {
	if discoveryUrl == nil {
		return fmt.Errorf("parameter discoveryUrl is required, but nil was provided")
	}

	if clientId == nil {
		return fmt.Errorf("parameter clientId is required, but nil was provided")
	}

	return nil
}

