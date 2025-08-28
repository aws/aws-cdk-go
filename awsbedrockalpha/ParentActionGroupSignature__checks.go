//go:build !no_runtime_type_checking

package awsbedrockalpha

import (
	"fmt"
)

func validateNewParentActionGroupSignatureParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

