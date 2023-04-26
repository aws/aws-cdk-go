//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"
)

func validateDockerBuildSecret_FromSrcParameters(src *string) error {
	if src == nil {
		return fmt.Errorf("parameter src is required, but nil was provided")
	}

	return nil
}

