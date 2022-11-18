//go:build !no_runtime_type_checking

package awscognito

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateUserPoolIdentityProvider_FromProviderNameParameters(scope constructs.Construct, id *string, providerName *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if providerName == nil {
		return fmt.Errorf("parameter providerName is required, but nil was provided")
	}

	return nil
}

