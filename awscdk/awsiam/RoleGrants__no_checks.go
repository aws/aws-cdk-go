//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RoleGrants) validateAssumeRoleParameters(identity IPrincipal) error {
	return nil
}

func (r *jsiiProxy_RoleGrants) validatePassRoleParameters(identity IPrincipal) error {
	return nil
}

func validateRoleGrants_FromRoleParameters(role interfacesawsiam.IRoleRef) error {
	return nil
}

