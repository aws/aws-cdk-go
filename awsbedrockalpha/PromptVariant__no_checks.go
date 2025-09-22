//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func validatePromptVariant_AgentParameters(props *AgentPromptVariantProps) error {
	return nil
}

func validatePromptVariant_ChatParameters(props *ChatPromptVariantProps) error {
	return nil
}

func validatePromptVariant_TextParameters(props *TextPromptVariantProps) error {
	return nil
}

