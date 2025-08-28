//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func (i *jsiiProxy_IPolicyValidationPluginBeta1) validateValidateParameters(context IPolicyValidationContextBeta1) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

