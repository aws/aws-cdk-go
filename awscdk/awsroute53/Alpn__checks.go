//go:build !no_runtime_type_checking

package awsroute53

import (
	"fmt"
)

func validateAlpn_OfParameters(protocol *string) error {
	if protocol == nil {
		return fmt.Errorf("parameter protocol is required, but nil was provided")
	}

	return nil
}

