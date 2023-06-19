//go:build no_runtime_type_checking

package integtests

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssertionsProvider) validateAddPolicyStatementFromSdkCallParameters(service *string, api *string) error {
	return nil
}

func (a *jsiiProxy_AssertionsProvider) validateEncodeParameters(obj interface{}) error {
	return nil
}

func (a *jsiiProxy_AssertionsProvider) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_AssertionsProvider) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateAssertionsProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAssertionsProviderParameters(scope constructs.Construct, id *string) error {
	return nil
}

