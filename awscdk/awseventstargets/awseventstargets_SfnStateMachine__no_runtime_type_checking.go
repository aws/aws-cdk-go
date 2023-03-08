//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SfnStateMachine) validateBindParameters(_rule awsevents.IRule) error {
	return nil
}

func validateNewSfnStateMachineParameters(machine awsstepfunctions.IStateMachine, props *SfnStateMachineProps) error {
	return nil
}

