//go:build !no_runtime_type_checking

package awsapigatewayv2

import (
	"fmt"
)

func validatePayloadFormatVersion_CustomParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

