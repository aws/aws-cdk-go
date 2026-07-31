//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func validatePropagateTags_ExplicitParameters(tags *map[string]*string) error {
	if tags == nil {
		return fmt.Errorf("parameter tags is required, but nil was provided")
	}

	return nil
}

