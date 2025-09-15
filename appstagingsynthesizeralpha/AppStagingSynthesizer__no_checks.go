//go:build no_runtime_type_checking

package appstagingsynthesizeralpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppStagingSynthesizer) validateAddBootstrapVersionRuleParameters(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) error {
	return nil
}

func (a *jsiiProxy_AppStagingSynthesizer) validateAddDockerImageAssetParameters(_asset *awscdk.DockerImageAssetSource) error {
	return nil
}

func (a *jsiiProxy_AppStagingSynthesizer) validateAddFileAssetParameters(_asset *awscdk.FileAssetSource) error {
	return nil
}

func (a *jsiiProxy_AppStagingSynthesizer) validateBindParameters(_stack awscdk.Stack) error {
	return nil
}

func (a *jsiiProxy_AppStagingSynthesizer) validateCloudFormationLocationFromDockerImageAssetParameters(dest *cloudassemblyschema.DockerImageDestination) error {
	return nil
}

func (a *jsiiProxy_AppStagingSynthesizer) validateCloudFormationLocationFromFileAssetParameters(location *cloudassemblyschema.FileDestination) error {
	return nil
}

func (a *jsiiProxy_AppStagingSynthesizer) validateEmitArtifactParameters(session awscdk.ISynthesisSession, options *awscdk.SynthesizeStackArtifactOptions) error {
	return nil
}

func (a *jsiiProxy_AppStagingSynthesizer) validateEmitStackArtifactParameters(stack awscdk.Stack, session awscdk.ISynthesisSession, options *awscdk.SynthesizeStackArtifactOptions) error {
	return nil
}

func (a *jsiiProxy_AppStagingSynthesizer) validateReusableBindParameters(stack awscdk.Stack) error {
	return nil
}

func (a *jsiiProxy_AppStagingSynthesizer) validateSynthesizeParameters(_session awscdk.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_AppStagingSynthesizer) validateSynthesizeStackTemplateParameters(stack awscdk.Stack, session awscdk.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_AppStagingSynthesizer) validateSynthesizeTemplateParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateAppStagingSynthesizer_CustomFactoryParameters(options *CustomFactoryOptions) error {
	return nil
}

func validateAppStagingSynthesizer_CustomResourcesParameters(options *CustomResourcesOptions) error {
	return nil
}

func validateAppStagingSynthesizer_DefaultResourcesParameters(options *DefaultResourcesOptions) error {
	return nil
}

