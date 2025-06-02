//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func (i *jsiiProxy_IListProducer) validateProduceParameters(context IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

