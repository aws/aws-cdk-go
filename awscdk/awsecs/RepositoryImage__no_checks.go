//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RepositoryImage) validateBindParameters(scope constructs.Construct, containerDefinition ContainerDefinition) error {
	return nil
}

func validateRepositoryImage_FromAssetParameters(directory *string, props *AssetImageProps) error {
	return nil
}

func validateRepositoryImage_FromDockerImageAssetParameters(asset awsecrassets.DockerImageAsset) error {
	return nil
}

func validateRepositoryImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

func validateRepositoryImage_FromRegistryParameters(name *string, props *RepositoryImageProps) error {
	return nil
}

func validateRepositoryImage_FromTarballParameters(tarballFile *string) error {
	return nil
}

func validateNewRepositoryImageParameters(imageName *string, props *RepositoryImageProps) error {
	return nil
}

