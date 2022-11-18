//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetImage) validateBindParameters(scope constructs.Construct, containerDefinition ContainerDefinition) error {
	return nil
}

func validateAssetImage_FromAssetParameters(directory *string, props *AssetImageProps) error {
	return nil
}

func validateAssetImage_FromDockerImageAssetParameters(asset awsecrassets.DockerImageAsset) error {
	return nil
}

func validateAssetImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

func validateAssetImage_FromRegistryParameters(name *string, props *RepositoryImageProps) error {
	return nil
}

func validateAssetImage_FromTarballParameters(tarballFile *string) error {
	return nil
}

func validateNewAssetImageParameters(directory *string, props *AssetImageProps) error {
	return nil
}

