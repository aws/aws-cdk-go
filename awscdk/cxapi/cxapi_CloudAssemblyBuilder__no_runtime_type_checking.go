//go:build no_runtime_type_checking

package cxapi

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudAssemblyBuilder) validateAddArtifactParameters(id *string, manifest *cloudassemblyschema.ArtifactManifest) error {
	return nil
}

func (c *jsiiProxy_CloudAssemblyBuilder) validateAddMissingParameters(missing *cloudassemblyschema.MissingContext) error {
	return nil
}

func (c *jsiiProxy_CloudAssemblyBuilder) validateBuildAssemblyParameters(options *AssemblyBuildOptions) error {
	return nil
}

func (c *jsiiProxy_CloudAssemblyBuilder) validateCreateNestedAssemblyParameters(artifactId *string, displayName *string) error {
	return nil
}

func validateNewCloudAssemblyBuilderParameters(props *CloudAssemblyBuilderProps) error {
	return nil
}

