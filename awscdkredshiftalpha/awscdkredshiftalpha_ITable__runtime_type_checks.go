//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha

import (
	"fmt"
)

func (i *jsiiProxy_ITable) validateGrantParameters(user IUser) error {
	if user == nil {
		return fmt.Errorf("parameter user is required, but nil was provided")
	}

	return nil
}

