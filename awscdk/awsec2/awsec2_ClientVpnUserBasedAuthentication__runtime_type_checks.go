//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

func validateClientVpnUserBasedAuthentication_ActiveDirectoryParameters(directoryId *string) error {
	if directoryId == nil {
		return fmt.Errorf("parameter directoryId is required, but nil was provided")
	}

	return nil
}

func validateClientVpnUserBasedAuthentication_FederatedParameters(samlProvider awsiam.ISamlProvider) error {
	if samlProvider == nil {
		return fmt.Errorf("parameter samlProvider is required, but nil was provided")
	}

	return nil
}

