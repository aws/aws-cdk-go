//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"
)

func validateNewReleaseLabelParameters(label *string) error {
	if label == nil {
		return fmt.Errorf("parameter label is required, but nil was provided")
	}

	return nil
}

