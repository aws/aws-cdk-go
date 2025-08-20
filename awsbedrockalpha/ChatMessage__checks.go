//go:build !no_runtime_type_checking

package awsbedrockalpha

import (
	"fmt"
)

func validateChatMessage_AssistantParameters(text *string) error {
	if text == nil {
		return fmt.Errorf("parameter text is required, but nil was provided")
	}

	return nil
}

func validateChatMessage_UserParameters(text *string) error {
	if text == nil {
		return fmt.Errorf("parameter text is required, but nil was provided")
	}

	return nil
}

func validateNewChatMessageParameters(role ChatMessageRole, text *string) error {
	if role == "" {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	if text == nil {
		return fmt.Errorf("parameter text is required, but nil was provided")
	}

	return nil
}

