//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BootstraplessSynthesizer) validateAddBootstrapVersionRuleParameters(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) error {
	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateAddDockerImageAssetParameters(_asset *DockerImageAssetSource) error {
	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateAddFileAssetParameters(_asset *FileAssetSource) error {
	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateBindParameters(stack Stack) error {
	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateCloudFormationLocationFromDockerImageAssetParameters(dest *cloudassemblyschema.DockerImageDestination) error {
	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateCloudFormationLocationFromFileAssetParameters(location *cloudassemblyschema.FileDestination) error {
	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateEmitArtifactParameters(session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateEmitStackArtifactParameters(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateSynthesizeStackTemplateParameters(stack Stack, session ISynthesisSession) error {
	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateSynthesizeTemplateParameters(session ISynthesisSession) error {
	return nil
}

func validateNewBootstraplessSynthesizerParameters(props *BootstraplessSynthesizerProps) error {
	return nil
}

