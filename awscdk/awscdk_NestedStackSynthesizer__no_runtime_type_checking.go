//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NestedStackSynthesizer) validateAddDockerImageAssetParameters(asset *DockerImageAssetSource) error {
	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateAddFileAssetParameters(asset *FileAssetSource) error {
	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateBindParameters(stack Stack) error {
	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateEmitStackArtifactParameters(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateSynthesizeStackTemplateParameters(stack Stack, session ISynthesisSession) error {
	return nil
}

func validateNewNestedStackSynthesizerParameters(parentDeployment IStackSynthesizer) error {
	return nil
}

