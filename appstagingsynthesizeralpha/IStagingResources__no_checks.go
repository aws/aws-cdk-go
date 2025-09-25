//go:build no_runtime_type_checking

package appstagingsynthesizeralpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IStagingResources) validateAddDockerImageParameters(asset *awscdk.DockerImageAssetSource) error {
	return nil
}

func (i *jsiiProxy_IStagingResources) validateAddFileParameters(asset *awscdk.FileAssetSource) error {
	return nil
}

