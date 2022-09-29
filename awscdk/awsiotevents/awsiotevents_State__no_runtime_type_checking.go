//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsiotevents

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_State) validateTransitionToParameters(targetState State, options *TransitionOptions) error {
	return nil
}

func validateNewStateParameters(props *StateProps) error {
	return nil
}

