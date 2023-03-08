//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StartExecution) validateBindParameters(task awsstepfunctions.Task) error {
	return nil
}

func validateNewStartExecutionParameters(stateMachine awsstepfunctions.IStateMachine, props *StartExecutionProps) error {
	return nil
}

