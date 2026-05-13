//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func validateArrayMergeStrategy_ReplaceByKeyParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

