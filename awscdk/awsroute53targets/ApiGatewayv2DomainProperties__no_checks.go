//go:build no_runtime_type_checking

package awsroute53targets

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiGatewayv2DomainProperties) validateBindParameters(_record awsroute53.IRecordSet) error {
	return nil
}

func validateNewApiGatewayv2DomainPropertiesParameters(regionalDomainName *string, regionalHostedZoneId *string) error {
	return nil
}

