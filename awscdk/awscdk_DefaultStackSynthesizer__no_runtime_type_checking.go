//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DefaultStackSynthesizer) validateAddDockerImageAssetParameters(asset *DockerImageAssetSource) error {
	return nil
}

func (d *jsiiProxy_DefaultStackSynthesizer) validateAddFileAssetParameters(asset *FileAssetSource) error {
	return nil
}

func (d *jsiiProxy_DefaultStackSynthesizer) validateBindParameters(stack Stack) error {
	return nil
}

func (d *jsiiProxy_DefaultStackSynthesizer) validateEmitStackArtifactParameters(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
	return nil
}

func (d *jsiiProxy_DefaultStackSynthesizer) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func (d *jsiiProxy_DefaultStackSynthesizer) validateSynthesizeStackTemplateParameters(stack Stack, session ISynthesisSession) error {
	return nil
}

func validateNewDefaultStackSynthesizerParameters(props *DefaultStackSynthesizerProps) error {
	return nil
}

