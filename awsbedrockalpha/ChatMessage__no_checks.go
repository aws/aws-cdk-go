//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateChatMessage_AssistantParameters(text *string) error {
	return nil
}

func validateChatMessage_UserParameters(text *string) error {
	return nil
}

func validateNewChatMessageParameters(role ChatMessageRole, text *string) error {
	return nil
}

