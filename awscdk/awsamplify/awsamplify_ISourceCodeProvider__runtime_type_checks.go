//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsamplify

import (
	"fmt"
)

func (i *jsiiProxy_ISourceCodeProvider) validateBindParameters(app App) error {
	if app == nil {
		return fmt.Errorf("parameter app is required, but nil was provided")
	}

	return nil
}

