//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"
)

func validateAuroraEngineVersion_OfParameters(auroraFullVersion *string) error {
	if auroraFullVersion == nil {
		return fmt.Errorf("parameter auroraFullVersion is required, but nil was provided")
	}

	return nil
}

