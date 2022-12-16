//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TagParameterContainerImage) validateBindParameters(scope constructs.Construct, containerDefinition ContainerDefinition) error {
	return nil
}

func validateTagParameterContainerImage_FromAssetParameters(directory *string, props *AssetImageProps) error {
	return nil
}

func validateTagParameterContainerImage_FromDockerImageAssetParameters(asset awsecrassets.DockerImageAsset) error {
	return nil
}

func validateTagParameterContainerImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

func validateTagParameterContainerImage_FromRegistryParameters(name *string, props *RepositoryImageProps) error {
	return nil
}

func validateTagParameterContainerImage_FromTarballParameters(tarballFile *string) error {
	return nil
}

func validateNewTagParameterContainerImageParameters(repository awsecr.IRepository) error {
	return nil
}

