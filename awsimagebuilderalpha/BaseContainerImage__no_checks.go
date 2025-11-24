//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateBaseContainerImage_FromDockerHubParameters(repository *string, tag *string) error {
	return nil
}

func validateBaseContainerImage_FromEcrParameters(repository awsecr.IRepository, tag *string) error {
	return nil
}

func validateBaseContainerImage_FromEcrPublicParameters(registryAlias *string, repositoryName *string, tag *string) error {
	return nil
}

func validateBaseContainerImage_FromStringParameters(baseContainerImageString *string) error {
	return nil
}

func validateNewBaseContainerImageParameters(image *string) error {
	return nil
}

