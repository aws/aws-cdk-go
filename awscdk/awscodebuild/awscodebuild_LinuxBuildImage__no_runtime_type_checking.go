//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LinuxBuildImage) validateRunScriptBuildspecParameters(entrypoint *string) error {
	return nil
}

func (l *jsiiProxy_LinuxBuildImage) validateValidateParameters(_arg *BuildEnvironment) error {
	return nil
}

func validateLinuxBuildImage_FromAssetParameters(scope constructs.Construct, id *string, props *awsecrassets.DockerImageAssetProps) error {
	return nil
}

func validateLinuxBuildImage_FromCodeBuildImageIdParameters(id *string) error {
	return nil
}

func validateLinuxBuildImage_FromDockerRegistryParameters(name *string, options *DockerImageOptions) error {
	return nil
}

func validateLinuxBuildImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

