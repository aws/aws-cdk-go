//go:build !no_runtime_type_checking

package awscognito

import (
	"fmt"
)

func validateOAuthScope_CustomParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateOAuthScope_ResourceServerParameters(server IUserPoolResourceServer, scope ResourceServerScope) error {
	if server == nil {
		return fmt.Errorf("parameter server is required, but nil was provided")
	}

	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

