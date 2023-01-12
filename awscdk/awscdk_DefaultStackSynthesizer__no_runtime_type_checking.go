//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DefaultStackSynthesizer) validateAddBootstrapVersionRuleParameters(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) error {
	return nil
}

func (d *jsiiProxy_DefaultStackSynthesizer) validateAddDockerImageAssetParameters(asset *DockerImageAssetSource) error {
	return nil
}

func (d *jsiiProxy_DefaultStackSynthesizer) validateAddFileAssetParameters(asset *FileAssetSource) error {
	return nil
}

func (d *jsiiProxy_DefaultStackSynthesizer) validateBindParameters(stack Stack) error {
	return nil
}

func (d *jsiiProxy_DefaultStackSynthesizer) validateCloudFormationLocationFromDockerImageAssetParameters(dest *cloudassemblyschema.DockerImageDestination) error {
	return nil
}

func (d *jsiiProxy_DefaultStackSynthesizer) validateCloudFormationLocationFromFileAssetParameters(location *cloudassemblyschema.FileDestination) error {
	return nil
}

func (d *jsiiProxy_DefaultStackSynthesizer) validateEmitArtifactParameters(session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
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

func (d *jsiiProxy_DefaultStackSynthesizer) validateSynthesizeTemplateParameters(session ISynthesisSession) error {
	return nil
}

func validateNewDefaultStackSynthesizerParameters(props *DefaultStackSynthesizerProps) error {
	return nil
}

