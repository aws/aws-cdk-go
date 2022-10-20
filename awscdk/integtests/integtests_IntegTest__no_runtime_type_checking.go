//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package integtests

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegTest) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (i *jsiiProxy_IntegTest) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateIntegTest_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewIntegTestParameters(scope constructs.Construct, id *string, props *IntegTestProps) error {
	return nil
}

