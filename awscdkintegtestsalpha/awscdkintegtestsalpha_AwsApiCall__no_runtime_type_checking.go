//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// CDK Integration Testing Constructs
package awscdkintegtestsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsApiCall) validateAssertAtPathParameters(path *string, expected ExpectedResult) error {
	return nil
}

func (a *jsiiProxy_AwsApiCall) validateExpectParameters(expected ExpectedResult) error {
	return nil
}

func (a *jsiiProxy_AwsApiCall) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsApiCall) validateGetAttStringParameters(attributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsApiCall) validateNextParameters(next IApiCall) error {
	return nil
}

func validateAwsApiCall_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsApiCall) validateSetFlattenResponseParameters(val *string) error {
	return nil
}

func validateNewAwsApiCallParameters(scope constructs.Construct, id *string, props *AwsApiCallProps) error {
	return nil
}

