//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LegacyStackSynthesizer) validateAddDockerImageAssetParameters(asset *DockerImageAssetSource) error {
	return nil
}

func (l *jsiiProxy_LegacyStackSynthesizer) validateAddFileAssetParameters(asset *FileAssetSource) error {
	return nil
}

func (l *jsiiProxy_LegacyStackSynthesizer) validateBindParameters(stack Stack) error {
	return nil
}

func (l *jsiiProxy_LegacyStackSynthesizer) validateEmitStackArtifactParameters(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
	return nil
}

func (l *jsiiProxy_LegacyStackSynthesizer) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func (l *jsiiProxy_LegacyStackSynthesizer) validateSynthesizeStackTemplateParameters(stack Stack, session ISynthesisSession) error {
	return nil
}

