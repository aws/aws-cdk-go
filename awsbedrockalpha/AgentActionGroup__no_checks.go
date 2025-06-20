//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateAgentActionGroup_CodeInterpreterParameters(enabled *bool) error {
	return nil
}

func validateAgentActionGroup_UserInputParameters(enabled *bool) error {
	return nil
}

func validateNewAgentActionGroupParameters(props *AgentActionGroupProps) error {
	return nil
}

