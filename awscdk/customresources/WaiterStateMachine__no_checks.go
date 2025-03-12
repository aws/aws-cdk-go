//go:build no_runtime_type_checking

package customresources

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WaiterStateMachine) validateGrantStartExecutionParameters(identity awsiam.IGrantable) error {
	return nil
}

func validateWaiterStateMachine_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewWaiterStateMachineParameters(scope constructs.Construct, id *string, props *WaiterStateMachineProps) error {
	return nil
}

