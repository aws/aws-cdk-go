//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func (i *jsiiProxy_IUser) validateAddToGroupParameters(group IGroup) error {
	if group == nil {
		return fmt.Errorf("parameter group is required, but nil was provided")
	}

	return nil
}

