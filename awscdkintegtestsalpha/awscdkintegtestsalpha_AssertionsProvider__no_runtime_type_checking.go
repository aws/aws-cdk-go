//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// CDK Integration Testing Constructs
package awscdkintegtestsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssertionsProvider) validateAddPolicyStatementFromSdkCallParameters(service *string, api *string) error {
	return nil
}

func (a *jsiiProxy_AssertionsProvider) validateAddToRolePolicyParameters(statement interface{}) error {
	return nil
}

func (a *jsiiProxy_AssertionsProvider) validateEncodeParameters(obj interface{}) error {
	return nil
}

func validateAssertionsProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAssertionsProviderParameters(scope constructs.Construct, id *string) error {
	return nil
}

