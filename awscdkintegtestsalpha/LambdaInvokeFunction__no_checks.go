//go:build no_runtime_type_checking

package awscdkintegtestsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaInvokeFunction) validateAssertAtPathParameters(path *string, expected ExpectedResult) error {
	return nil
}

func (l *jsiiProxy_LambdaInvokeFunction) validateExpectParameters(expected ExpectedResult) error {
	return nil
}

func (l *jsiiProxy_LambdaInvokeFunction) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (l *jsiiProxy_LambdaInvokeFunction) validateGetAttStringParameters(attributeName *string) error {
	return nil
}

func (l *jsiiProxy_LambdaInvokeFunction) validateNextParameters(next IApiCall) error {
	return nil
}

func (l *jsiiProxy_LambdaInvokeFunction) validateWaitForAssertionsParameters(options *WaiterStateMachineOptions) error {
	return nil
}

func validateLambdaInvokeFunction_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_LambdaInvokeFunction) validateSetFlattenResponseParameters(val *string) error {
	return nil
}

func validateNewLambdaInvokeFunctionParameters(scope constructs.Construct, id *string, props *LambdaInvokeFunctionProps) error {
	return nil
}

