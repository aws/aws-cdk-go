//go:build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"
)

func (i *jsiiProxy_INextable) validateNextParameters(state IChainable) error {
	if state == nil {
		return fmt.Errorf("parameter state is required, but nil was provided")
	}

	return nil
}

