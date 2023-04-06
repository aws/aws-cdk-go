//go:build !no_runtime_type_checking

package awsdocdb

import (
	"fmt"
)

func validateNewEndpointParameters(address *string, port *float64) error {
	if address == nil {
		return fmt.Errorf("parameter address is required, but nil was provided")
	}

	if port == nil {
		return fmt.Errorf("parameter port is required, but nil was provided")
	}

	return nil
}

