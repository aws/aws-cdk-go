//go:build no_runtime_type_checking

package assertions

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Annotations) validateFindErrorParameters(constructPath *string, message interface{}) error {
	return nil
}

func (a *jsiiProxy_Annotations) validateFindInfoParameters(constructPath *string, message interface{}) error {
	return nil
}

func (a *jsiiProxy_Annotations) validateFindWarningParameters(constructPath *string, message interface{}) error {
	return nil
}

func (a *jsiiProxy_Annotations) validateHasErrorParameters(constructPath *string, message interface{}) error {
	return nil
}

func (a *jsiiProxy_Annotations) validateHasInfoParameters(constructPath *string, message interface{}) error {
	return nil
}

func (a *jsiiProxy_Annotations) validateHasNoErrorParameters(constructPath *string, message interface{}) error {
	return nil
}

func (a *jsiiProxy_Annotations) validateHasNoInfoParameters(constructPath *string, message interface{}) error {
	return nil
}

func (a *jsiiProxy_Annotations) validateHasNoWarningParameters(constructPath *string, message interface{}) error {
	return nil
}

func (a *jsiiProxy_Annotations) validateHasWarningParameters(constructPath *string, message interface{}) error {
	return nil
}

func validateAnnotations_FromStackParameters(stack awscdk.Stack) error {
	return nil
}

