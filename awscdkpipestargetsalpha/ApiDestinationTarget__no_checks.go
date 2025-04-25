//go:build no_runtime_type_checking

package awscdkpipestargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiDestinationTarget) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (a *jsiiProxy_ApiDestinationTarget) validateGrantPushParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewApiDestinationTargetParameters(destination awsevents.IApiDestination, parameters *ApiDestinationTargetParameters) error {
	return nil
}

