//go:build no_runtime_type_checking

package awscodepipeline

// Building without runtime type checking enabled, so all the below just return nil

func validateNewTriggerParameters(props *TriggerProps) error {
	return nil
}

