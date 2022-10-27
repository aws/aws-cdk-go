//go:build !no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	"fmt"
)

func (i *jsiiProxy_IPostProcessor) validatePostProcessParameters(input interface{}, context IResolveContext) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}

	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

