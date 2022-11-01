//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LinuxArmBuildImage) validateRunScriptBuildspecParameters(entrypoint *string) error {
	return nil
}

func (l *jsiiProxy_LinuxArmBuildImage) validateValidateParameters(buildEnvironment *BuildEnvironment) error {
	return nil
}

func validateLinuxArmBuildImage_FromCodeBuildImageIdParameters(id *string) error {
	return nil
}

func validateLinuxArmBuildImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

