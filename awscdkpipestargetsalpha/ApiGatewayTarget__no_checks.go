//go:build no_runtime_type_checking

package awscdkpipestargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiGatewayTarget) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (a *jsiiProxy_ApiGatewayTarget) validateGrantPushParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewApiGatewayTargetParameters(restApi awsapigateway.IRestApi, parameters *ApiGatewayTargetParameters) error {
	return nil
}

