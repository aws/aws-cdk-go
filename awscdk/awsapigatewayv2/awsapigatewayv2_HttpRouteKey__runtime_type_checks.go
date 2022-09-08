//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsapigatewayv2

import (
	"fmt"
)

func validateHttpRouteKey_WithParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

