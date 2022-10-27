//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"
)

func validateNewFileSetParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

