//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcrImage) validateBindParameters(_scope constructs.Construct, containerDefinition ContainerDefinition) error {
	return nil
}

func validateEcrImage_FromAssetParameters(directory *string, props *AssetImageProps) error {
	return nil
}

func validateEcrImage_FromDockerImageAssetParameters(asset awsecrassets.DockerImageAsset) error {
	return nil
}

func validateEcrImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

func validateEcrImage_FromRegistryParameters(name *string, props *RepositoryImageProps) error {
	return nil
}

func validateEcrImage_FromTarballParameters(tarballFile *string) error {
	return nil
}

func validateNewEcrImageParameters(repository awsecr.IRepository, tagOrDigest *string) error {
	return nil
}

