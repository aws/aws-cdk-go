//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func (i *jsiiProxy_IArrayMergeStrategy) validateMergeParameters(target *[]interface{}, source *[]interface{}) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}

	return nil
}

