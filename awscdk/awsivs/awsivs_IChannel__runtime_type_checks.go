//go:build !no_runtime_type_checking

package awsivs

import (
	"fmt"
)

func (i *jsiiProxy_IChannel) validateAddStreamKeyParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

