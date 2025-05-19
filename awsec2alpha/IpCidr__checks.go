//go:build !no_runtime_type_checking

package awsec2alpha

import (
	"fmt"
)

func validateNewIpCidrParameters(props *string) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

