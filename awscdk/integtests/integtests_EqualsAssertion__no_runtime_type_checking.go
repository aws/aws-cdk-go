//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package integtests

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EqualsAssertion) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (e *jsiiProxy_EqualsAssertion) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateEqualsAssertion_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEqualsAssertionParameters(scope constructs.Construct, id *string, props *EqualsAssertionProps) error {
	return nil
}

