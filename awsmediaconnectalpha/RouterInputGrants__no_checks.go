//go:build no_runtime_type_checking

package awsmediaconnectalpha

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RouterInputGrants) validateActionsParameters(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) error {
	return nil
}

func (r *jsiiProxy_RouterInputGrants) validateRestartParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RouterInputGrants) validateStartParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RouterInputGrants) validateStopParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateRouterInputGrants_FromRouterInputParameters(resource interfacesawsmediaconnect.IRouterInputRef) error {
	return nil
}

