//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"
)

func validateSessionPinningFilter_OfParameters(filterName *string) error {
	if filterName == nil {
		return fmt.Errorf("parameter filterName is required, but nil was provided")
	}

	return nil
}

