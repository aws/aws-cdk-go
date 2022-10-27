//go:build !no_runtime_type_checking

package awsapigatewayv2

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

func validateIntegrationCredentials_FromRoleParameters(role awsiam.IRole) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

