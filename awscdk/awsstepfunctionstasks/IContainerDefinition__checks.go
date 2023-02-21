//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"
)

func (i *jsiiProxy_IContainerDefinition) validateBindParameters(task ISageMakerTask) error {
	if task == nil {
		return fmt.Errorf("parameter task is required, but nil was provided")
	}

	return nil
}

