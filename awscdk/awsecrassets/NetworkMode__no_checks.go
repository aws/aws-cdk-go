//go:build no_runtime_type_checking

package awsecrassets

// Building without runtime type checking enabled, so all the below just return nil

func validateNetworkMode_CustomParameters(mode *string) error {
	return nil
}

func validateNetworkMode_FromContainerParameters(containerId *string) error {
	return nil
}

