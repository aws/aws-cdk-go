//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MacBuildImage) validateRunScriptBuildspecParameters(entrypoint *string) error {
	return nil
}

func (m *jsiiProxy_MacBuildImage) validateValidateParameters(buildEnvironment *BuildEnvironment) error {
	return nil
}

func validateMacBuildImage_FromAssetParameters(scope constructs.Construct, id *string, props *awsecrassets.DockerImageAssetProps) error {
	return nil
}

func validateMacBuildImage_FromDockerRegistryParameters(name *string, options *DockerImageOptions) error {
	return nil
}

func validateMacBuildImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

