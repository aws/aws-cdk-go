//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func validatePolicyAction_AnyOfParameters(actions *[]*string) error {
	return nil
}

func validatePolicyAction_OneParameters(action *string) error {
	return nil
}

