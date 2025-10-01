//go:build !no_runtime_type_checking

package awscdkelasticachealpha

import (
	"fmt"
)

func (i *jsiiProxy_IUserGroup) validateAddUserParameters(user IUser) error {
	if user == nil {
		return fmt.Errorf("parameter user is required, but nil was provided")
	}

	return nil
}

