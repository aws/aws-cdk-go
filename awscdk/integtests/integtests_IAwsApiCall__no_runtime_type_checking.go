//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package integtests

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAwsApiCall) validateAssertAtPathParameters(path *string, expected ExpectedResult) error {
	return nil
}

func (i *jsiiProxy_IAwsApiCall) validateExpectParameters(expected ExpectedResult) error {
	return nil
}

func (i *jsiiProxy_IAwsApiCall) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (i *jsiiProxy_IAwsApiCall) validateGetAttStringParameters(attributeName *string) error {
	return nil
}

