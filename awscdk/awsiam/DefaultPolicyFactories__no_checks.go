//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func validateDefaultPolicyFactories_GetParameters(type_ *string) error {
	return nil
}

func validateDefaultPolicyFactories_HasParameters(type_ *string) error {
	return nil
}

func validateDefaultPolicyFactories_SetParameters(type_ *string, factory IResourcePolicyFactory) error {
	return nil
}

