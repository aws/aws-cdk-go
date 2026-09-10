//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func validatePolicyStatement_FromCedarParameters(cedarStatement *string) error {
	return nil
}

func validateNewPolicyStatementParameters(props *PolicyStatementProps) error {
	return nil
}

