//go:build no_runtime_type_checking

package integtests

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

func (a *jsiiProxy_AwsApiCall) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_AwsApiCall) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateAwsApiCall_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsApiCall) validateSetProviderParameters(val AssertionsProvider) error {
	return nil
}

func validateNewAwsApiCallParameters(scope constructs.Construct, id *string, props *AwsApiCallProps) error {
	return nil
}

