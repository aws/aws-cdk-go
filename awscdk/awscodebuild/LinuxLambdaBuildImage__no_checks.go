//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LinuxLambdaBuildImage) validateRunScriptBuildspecParameters(entrypoint *string) error {
	return nil
}

func (l *jsiiProxy_LinuxLambdaBuildImage) validateValidateParameters(buildEnvironment *BuildEnvironment) error {
	return nil
}

