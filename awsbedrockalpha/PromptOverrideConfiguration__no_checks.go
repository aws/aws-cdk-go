//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func validatePromptOverrideConfiguration_FromStepsParameters(steps *[]*PromptStepConfigBase) error {
	return nil
}

func validatePromptOverrideConfiguration_WithCustomParserParameters(props *CustomParserProps) error {
	return nil
}

