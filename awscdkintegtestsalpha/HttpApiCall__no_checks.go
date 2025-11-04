//go:build no_runtime_type_checking

package awscdkintegtestsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpApiCall) validateAssertAtPathParameters(path *string, expected ExpectedResult) error {
	return nil
}

func (h *jsiiProxy_HttpApiCall) validateExpectParameters(expected ExpectedResult) error {
	return nil
}

func (h *jsiiProxy_HttpApiCall) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (h *jsiiProxy_HttpApiCall) validateGetAttStringParameters(attributeName *string) error {
	return nil
}

func (h *jsiiProxy_HttpApiCall) validateNextParameters(next IApiCall) error {
	return nil
}

func (h *jsiiProxy_HttpApiCall) validateWaitForAssertionsParameters(options *WaiterStateMachineOptions) error {
	return nil
}

func validateHttpApiCall_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_HttpApiCall) validateSetFlattenResponseParameters(val *string) error {
	return nil
}

func validateNewHttpApiCallParameters(scope constructs.Construct, id *string, props *HttpCallProps) error {
	return nil
}

