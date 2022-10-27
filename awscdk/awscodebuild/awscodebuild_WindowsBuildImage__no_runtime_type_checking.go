//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WindowsBuildImage) validateRunScriptBuildspecParameters(entrypoint *string) error {
	return nil
}

func (w *jsiiProxy_WindowsBuildImage) validateValidateParameters(buildEnvironment *BuildEnvironment) error {
	return nil
}

func validateWindowsBuildImage_FromAssetParameters(scope constructs.Construct, id *string, props *awsecrassets.DockerImageAssetProps) error {
	return nil
}

func validateWindowsBuildImage_FromDockerRegistryParameters(name *string, options *DockerImageOptions) error {
	return nil
}

func validateWindowsBuildImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

