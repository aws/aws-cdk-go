//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StackSynthesizer) validateAddDockerImageAssetParameters(asset *DockerImageAssetSource) error {
	return nil
}

func (s *jsiiProxy_StackSynthesizer) validateAddFileAssetParameters(asset *FileAssetSource) error {
	return nil
}

func (s *jsiiProxy_StackSynthesizer) validateBindParameters(stack Stack) error {
	return nil
}

func (s *jsiiProxy_StackSynthesizer) validateEmitStackArtifactParameters(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
	return nil
}

func (s *jsiiProxy_StackSynthesizer) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_StackSynthesizer) validateSynthesizeStackTemplateParameters(stack Stack, session ISynthesisSession) error {
	return nil
}

