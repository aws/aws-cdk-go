//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetManifestBuilder) validateAddDockerImageAssetParameters(stack Stack, sourceHash *string, source *cloudassemblyschema.DockerImageSource, dest *cloudassemblyschema.DockerImageDestination, options *AddDockerImageAssetOptions) error {
	return nil
}

func (a *jsiiProxy_AssetManifestBuilder) validateAddFileAssetParameters(stack Stack, sourceHash *string, source *cloudassemblyschema.FileSource, dest *cloudassemblyschema.FileDestination, options *AddFileAssetOptions) error {
	return nil
}

func (a *jsiiProxy_AssetManifestBuilder) validateDefaultAddDockerImageAssetParameters(stack Stack, asset *DockerImageAssetSource, target *AssetManifestDockerImageDestination, options *AddDockerImageAssetOptions) error {
	return nil
}

func (a *jsiiProxy_AssetManifestBuilder) validateDefaultAddFileAssetParameters(stack Stack, asset *FileAssetSource, target *AssetManifestFileDestination, options *AddFileAssetOptions) error {
	return nil
}

func (a *jsiiProxy_AssetManifestBuilder) validateEmitManifestParameters(stack Stack, session ISynthesisSession, options *cloudassemblyschema.AssetManifestOptions) error {
	return nil
}

