//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func validateEncryptedResources_OfParameters(resource interfaces.IEnvironmentAware) error {
	return nil
}

func validateEncryptedResources_RegisterParameters(scope constructs.IConstruct, cfnType *string, factory IEncryptedResourceFactory) error {
	return nil
}

