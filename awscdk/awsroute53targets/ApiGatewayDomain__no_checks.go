//go:build no_runtime_type_checking

package awsroute53targets

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiGatewayDomain) validateBindParameters(_record awsroute53.IRecordSet) error {
	return nil
}

func validateNewApiGatewayDomainParameters(domainName awsapigateway.IDomainName) error {
	return nil
}

