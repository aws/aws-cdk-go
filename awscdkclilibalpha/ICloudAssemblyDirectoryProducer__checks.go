//go:build !no_runtime_type_checking

package awscdkclilibalpha

import (
	"fmt"
)

func (i *jsiiProxy_ICloudAssemblyDirectoryProducer) validateProduceParameters(context *map[string]interface{}) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

