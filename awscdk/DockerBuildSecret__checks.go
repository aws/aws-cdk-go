//go:build !no_runtime_type_checking

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

