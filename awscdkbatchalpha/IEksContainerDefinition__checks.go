//go:build !no_runtime_type_checking

package awscdkbatchalpha

import (
	"fmt"
)

func (i *jsiiProxy_IEksContainerDefinition) validateAddVolumeParameters(volume EksVolume) error {
	if volume == nil {
		return fmt.Errorf("parameter volume is required, but nil was provided")
	}

	return nil
}

