//go:build !no_runtime_type_checking

package awslogs

import (
	"fmt"
)

func validateNewJsonPatternParameters(jsonPatternString *string) error {
	if jsonPatternString == nil {
		return fmt.Errorf("parameter jsonPatternString is required, but nil was provided")
	}

	return nil
}

