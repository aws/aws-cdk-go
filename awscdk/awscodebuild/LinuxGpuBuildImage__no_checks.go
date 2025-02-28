//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LinuxGpuBuildImage) validateBindParameters(scope constructs.Construct, project IProject, _options *BuildImageBindOptions) error {
	return nil
}

func (l *jsiiProxy_LinuxGpuBuildImage) validateRunScriptBuildspecParameters(entrypoint *string) error {
	return nil
}

func (l *jsiiProxy_LinuxGpuBuildImage) validateValidateParameters(buildEnvironment *BuildEnvironment) error {
	return nil
}

func validateLinuxGpuBuildImage_AwsDeepLearningContainersImageParameters(repositoryName *string, tag *string) error {
	return nil
}

func validateLinuxGpuBuildImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

