//go:build !no_runtime_type_checking

package awsmediaconnectalpha

import (
	"fmt"
)

func validateSourcePriorityConfig_PrimarySecondaryParameters(primary PrimarySource) error {
	if primary == "" {
		return fmt.Errorf("parameter primary is required, but nil was provided")
	}

	return nil
}

