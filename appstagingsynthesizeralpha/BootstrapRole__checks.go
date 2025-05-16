//go:build !no_runtime_type_checking

package appstagingsynthesizeralpha

import (
	"fmt"
)

func validateBootstrapRole_FromRoleArnParameters(arn *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

