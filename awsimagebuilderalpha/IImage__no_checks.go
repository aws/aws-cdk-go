//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IImage) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IImage) validateGrantDefaultExecutionRolePermissionsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IImage) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

