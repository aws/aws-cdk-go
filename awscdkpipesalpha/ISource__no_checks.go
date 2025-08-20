//go:build no_runtime_type_checking

package awscdkpipesalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ISource) validateBindParameters(pipe IPipe) error {
	return nil
}

func (i *jsiiProxy_ISource) validateGrantReadParameters(grantee awsiam.IRole) error {
	return nil
}

