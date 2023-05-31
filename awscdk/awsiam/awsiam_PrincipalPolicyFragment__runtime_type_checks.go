//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func validateNewPrincipalPolicyFragmentParameters(principalJson *map[string]*[]*string) error {
	if principalJson == nil {
		return fmt.Errorf("parameter principalJson is required, but nil was provided")
	}

	return nil
}

