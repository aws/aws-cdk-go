//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	"fmt"
)

func validateNewOutputParameters(outputName *string) error {
	if outputName == nil {
		return fmt.Errorf("parameter outputName is required, but nil was provided")
	}

	return nil
}

