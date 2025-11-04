//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
)

func validateRuntimeAuthorizerConfiguration_UsingCognitoParameters(userPool awscognito.IUserPool, userPoolClients *[]awscognito.IUserPoolClient) error {
	if userPool == nil {
		return fmt.Errorf("parameter userPool is required, but nil was provided")
	}

	if userPoolClients == nil {
		return fmt.Errorf("parameter userPoolClients is required, but nil was provided")
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

