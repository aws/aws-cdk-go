//go:build no_runtime_type_checking

package triggers

// Building without runtime type checking enabled, so all the below just return nil

func validateTrigger_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTriggerParameters(scope constructs.Construct, id *string, props *TriggerProps) error {
	return nil
}

