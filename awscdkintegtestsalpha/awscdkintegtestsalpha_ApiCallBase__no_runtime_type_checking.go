//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// CDK Integration Testing Constructs
package awscdkintegtestsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiCallBase) validateAssertAtPathParameters(path *string, expected ExpectedResult) error {
	return nil
}

func (a *jsiiProxy_ApiCallBase) validateExpectParameters(expected ExpectedResult) error {
	return nil
}

func (a *jsiiProxy_ApiCallBase) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApiCallBase) validateGetAttStringParameters(attributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApiCallBase) validateNextParameters(next IApiCall) error {
	return nil
}

func (a *jsiiProxy_ApiCallBase) validateWaitForAssertionsParameters(options *WaiterStateMachineOptions) error {
	return nil
}

func validateApiCallBase_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiCallBase) validateSetFlattenResponseParameters(val *string) error {
	return nil
}

func validateNewApiCallBaseParameters(scope constructs.Construct, id *string) error {
	return nil
}

