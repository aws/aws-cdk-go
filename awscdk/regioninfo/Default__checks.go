//go:build !no_runtime_type_checking

package regioninfo

import (
	"fmt"
)

func validateDefault_ServicePrincipalParameters(serviceFqn *string, region *string, urlSuffix *string) error {
	if serviceFqn == nil {
		return fmt.Errorf("parameter serviceFqn is required, but nil was provided")
	}

	if region == nil {
		return fmt.Errorf("parameter region is required, but nil was provided")
	}

	if urlSuffix == nil {
		return fmt.Errorf("parameter urlSuffix is required, but nil was provided")
	}

	return nil
}

