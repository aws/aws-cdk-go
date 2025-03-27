//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerImage) validateBindParameters(scope constructs.Construct, containerDefinition ContainerDefinition) error {
	return nil
}

func validateContainerImage_FromAssetParameters(directory *string, props *AssetImageProps) error {
	return nil
}

func validateContainerImage_FromDockerImageAssetParameters(asset awsecrassets.DockerImageAsset) error {
	return nil
}

func validateContainerImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

func validateContainerImage_FromRegistryParameters(name *string, props *RepositoryImageProps) error {
	return nil
}

func validateContainerImage_FromTarballParameters(tarballFile *string) error {
	return nil
}

