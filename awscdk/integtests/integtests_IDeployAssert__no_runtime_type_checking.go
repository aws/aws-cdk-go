//go:build no_runtime_type_checking

package integtests

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDeployAssert) validateAwsApiCallParameters(service *string, api *string) error {
	return nil
}

func (i *jsiiProxy_IDeployAssert) validateExpectParameters(id *string, expected ExpectedResult, actual ActualResult) error {
	return nil
}

func (i *jsiiProxy_IDeployAssert) validateInvokeFunctionParameters(props *LambdaInvokeFunctionProps) error {
	return nil
}

