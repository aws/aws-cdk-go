//go:build !no_runtime_type_checking

package awscdkpipesalpha

import (
	"fmt"
)

func (i *jsiiProxy_IInputTransformation) validateBindParameters(pipe IPipe) error {
	if pipe == nil {
		return fmt.Errorf("parameter pipe is required, but nil was provided")
	}

	return nil
}

