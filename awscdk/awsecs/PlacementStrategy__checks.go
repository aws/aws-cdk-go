//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"
)

func validatePlacementStrategy_PackedByParameters(resource BinPackResource) error {
	if resource == "" {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

