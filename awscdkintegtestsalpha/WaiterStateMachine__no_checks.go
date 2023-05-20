//go:build no_runtime_type_checking

package awscdkintegtestsalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateWaiterStateMachine_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewWaiterStateMachineParameters(scope constructs.Construct, id *string, props *WaiterStateMachineProps) error {
	return nil
}

