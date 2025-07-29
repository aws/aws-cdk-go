//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func (i *jsiiProxy_IInspectable) validateInspectParameters(inspector TreeInspector) error {
	if inspector == nil {
		return fmt.Errorf("parameter inspector is required, but nil was provided")
	}

	return nil
}

