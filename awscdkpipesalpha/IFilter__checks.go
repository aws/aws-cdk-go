//go:build !no_runtime_type_checking

package awscdkpipesalpha

import (
	"fmt"
)

func (j *jsiiProxy_IFilter) validateSetFiltersParameters(val *[]IFilterPattern) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

