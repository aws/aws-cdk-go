//go:build no_runtime_type_checking

package integtests

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegTestCase) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (i *jsiiProxy_IntegTestCase) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateIntegTestCase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewIntegTestCaseParameters(scope constructs.Construct, id *string, props *IntegTestCaseProps) error {
	return nil
}

