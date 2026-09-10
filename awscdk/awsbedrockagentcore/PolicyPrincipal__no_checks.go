//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func validatePolicyPrincipal_EntityParameters(entityType *string, entityId *string) error {
	return nil
}

func validatePolicyPrincipal_EntityTypeParameters(entityType *string) error {
	return nil
}

func validatePolicyPrincipal_InGroupParameters(groupType *string, groupId *string) error {
	return nil
}

