//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsstepfunctions

import (
	"fmt"
)

func (i *jsiiProxy_IStepFunctionsTask) validateBindParameters(task Task) error {
	if task == nil {
		return fmt.Errorf("parameter task is required, but nil was provided")
	}

	return nil
}

