//go:build no_runtime_type_checking

package awscdkpipesalpha

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamicInput) validateResolveParameters(_context awscdk.IResolveContext) error {
	return nil
}

func validateDynamicInput_FromEventPathParameters(path *string) error {
	return nil
}

