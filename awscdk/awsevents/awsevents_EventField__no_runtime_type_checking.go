//go:build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventField) validateResolveParameters(_ctx awscdk.IResolveContext) error {
	return nil
}

func validateEventField_FromPathParameters(path *string) error {
	return nil
}

