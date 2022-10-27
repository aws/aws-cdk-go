//go:build no_runtime_type_checking

package awsioteventsactions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SetVariableAction) validateBindParameters(_scope constructs.Construct, _options *awsiotevents.ActionBindOptions) error {
	return nil
}

func validateNewSetVariableActionParameters(variableName *string, value awsiotevents.Expression) error {
	return nil
}

