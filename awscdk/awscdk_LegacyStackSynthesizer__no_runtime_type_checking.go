//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LegacyStackSynthesizer) validateAddBootstrapVersionRuleParameters(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) error {
	return nil
}

func (l *jsiiProxy_LegacyStackSynthesizer) validateAddDockerImageAssetParameters(asset *DockerImageAssetSource) error {
	return nil
}

func (l *jsiiProxy_LegacyStackSynthesizer) validateAddFileAssetParameters(asset *FileAssetSource) error {
	return nil
}

func (l *jsiiProxy_LegacyStackSynthesizer) validateBindParameters(stack Stack) error {
	return nil
}

func (l *jsiiProxy_LegacyStackSynthesizer) validateCloudFormationLocationFromDockerImageAssetParameters(dest *cloudassemblyschema.DockerImageDestination) error {
	return nil
}

func (l *jsiiProxy_LegacyStackSynthesizer) validateCloudFormationLocationFromFileAssetParameters(location *cloudassemblyschema.FileDestination) error {
	return nil
}

func (l *jsiiProxy_LegacyStackSynthesizer) validateEmitArtifactParameters(session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
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

func (l *jsiiProxy_LegacyStackSynthesizer) validateSynthesizeTemplateParameters(session ISynthesisSession) error {
	return nil
}

