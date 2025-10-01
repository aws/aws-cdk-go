//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func validatePromptTemplateConfiguration_ChatParameters(props *ChatTemplateConfigurationProps) error {
	return nil
}

func validatePromptTemplateConfiguration_TextParameters(props *TextTemplateConfigurationProps) error {
	return nil
}

