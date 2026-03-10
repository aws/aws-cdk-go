//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func (i *jsiiProxy_IMergeStrategy) validateApplyParameters(target *map[string]interface{}, source *map[string]interface{}, allowedKeys *[]*string) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}

	if allowedKeys == nil {
		return fmt.Errorf("parameter allowedKeys is required, but nil was provided")
	}

	return nil
}

