//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func validatePolicyAttribute_ContextParameters(attribute *string) error {
	return nil
}

func validatePolicyAttribute_PrincipalParameters(attribute *string) error {
	return nil
}

func validatePolicyAttribute_ResourceParameters(attribute *string) error {
	return nil
}

