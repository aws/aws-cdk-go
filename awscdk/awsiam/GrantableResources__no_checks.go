//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func validateGrantableResources_IsEncryptedResourceParameters(resource constructs.IConstruct) error {
	return nil
}

func validateGrantableResources_IsResourceWithPolicyParameters(resource interfaces.IEnvironmentAware) error {
	return nil
}

