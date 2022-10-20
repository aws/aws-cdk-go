//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// CDK Integration Testing Constructs
package awscdkintegtestsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IApiCall) validateAssertAtPathParameters(path *string, expected ExpectedResult) error {
	return nil
}

func (i *jsiiProxy_IApiCall) validateExpectParameters(expected ExpectedResult) error {
	return nil
}

func (i *jsiiProxy_IApiCall) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (i *jsiiProxy_IApiCall) validateGetAttStringParameters(attributeName *string) error {
	return nil
}

func (i *jsiiProxy_IApiCall) validateNextParameters(next IApiCall) error {
	return nil
}

func (i *jsiiProxy_IApiCall) validateWaitForAssertionsParameters(options *WaiterStateMachineOptions) error {
	return nil
}

