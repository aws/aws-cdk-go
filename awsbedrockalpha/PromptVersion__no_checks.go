//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func validatePromptVersion_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPromptVersionParameters(scope constructs.Construct, id *string, props *PromptVersionProps) error {
	return nil
}

