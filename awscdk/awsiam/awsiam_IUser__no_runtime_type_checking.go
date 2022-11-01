//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IUser) validateAddToGroupParameters(group IGroup) error {
	return nil
}

