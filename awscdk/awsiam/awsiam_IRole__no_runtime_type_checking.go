//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IRole) validateGrantParameters(grantee IPrincipal) error {
	return nil
}

func (i *jsiiProxy_IRole) validateGrantAssumeRoleParameters(grantee IPrincipal) error {
	return nil
}

func (i *jsiiProxy_IRole) validateGrantPassRoleParameters(grantee IPrincipal) error {
	return nil
}

