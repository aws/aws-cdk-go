//go:build no_runtime_type_checking

package awscdkpipesenrichmentsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiGatewayEnrichment) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (a *jsiiProxy_ApiGatewayEnrichment) validateGrantInvokeParameters(pipeRole awsiam.IRole) error {
	return nil
}

func validateNewApiGatewayEnrichmentParameters(restApi awsapigateway.IRestApi, props *ApiGatewayEnrichmentProps) error {
	return nil
}

