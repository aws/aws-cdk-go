//go:build no_runtime_type_checking

package awsroute53targets

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiGateway) validateBindParameters(_record awsroute53.IRecordSet) error {
	return nil
}

func validateNewApiGatewayParameters(api awsapigateway.RestApiBase) error {
	return nil
}

