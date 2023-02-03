//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StackSynthesizer) validateAddBootstrapVersionRuleParameters(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) error {
	return nil
}

func (s *jsiiProxy_StackSynthesizer) validateAddDockerImageAssetParameters(asset *DockerImageAssetSource) error {
	return nil
}

func (s *jsiiProxy_StackSynthesizer) validateAddFileAssetParameters(asset *FileAssetSource) error {
	return nil
}

func (s *jsiiProxy_StackSynthesizer) validateBindParameters(stack Stack) error {
	return nil
}

func (s *jsiiProxy_StackSynthesizer) validateCloudFormationLocationFromDockerImageAssetParameters(dest *cloudassemblyschema.DockerImageDestination) error {
	return nil
}

func (s *jsiiProxy_StackSynthesizer) validateCloudFormationLocationFromFileAssetParameters(location *cloudassemblyschema.FileDestination) error {
	return nil
}

func (s *jsiiProxy_StackSynthesizer) validateEmitArtifactParameters(session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
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

func (s *jsiiProxy_StackSynthesizer) validateSynthesizeTemplateParameters(session ISynthesisSession) error {
	return nil
}

