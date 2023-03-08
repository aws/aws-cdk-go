//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NestedStackSynthesizer) validateAddBootstrapVersionRuleParameters(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) error {
	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateAddDockerImageAssetParameters(asset *DockerImageAssetSource) error {
	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateAddFileAssetParameters(asset *FileAssetSource) error {
	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateBindParameters(stack Stack) error {
	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateCloudFormationLocationFromDockerImageAssetParameters(dest *cloudassemblyschema.DockerImageDestination) error {
	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateCloudFormationLocationFromFileAssetParameters(location *cloudassemblyschema.FileDestination) error {
	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateEmitArtifactParameters(session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
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

func (n *jsiiProxy_NestedStackSynthesizer) validateSynthesizeTemplateParameters(session ISynthesisSession) error {
	return nil
}

func validateNewNestedStackSynthesizerParameters(parentDeployment IStackSynthesizer) error {
	return nil
}

