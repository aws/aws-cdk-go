//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsredshift

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITable) validateGrantParameters(user IUser) error {
	return nil
}

