//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func validatePolicyResource_AnyOfTypeParameters(entityType *string) error {
	return nil
}

func validatePolicyResource_InstanceParameters(entityType *string, entityArn *string) error {
	return nil
}

