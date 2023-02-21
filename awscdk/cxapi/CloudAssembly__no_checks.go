//go:build no_runtime_type_checking

package cxapi

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudAssembly) validateGetNestedAssemblyParameters(artifactId *string) error {
	return nil
}

func (c *jsiiProxy_CloudAssembly) validateGetNestedAssemblyArtifactParameters(artifactId *string) error {
	return nil
}

func (c *jsiiProxy_CloudAssembly) validateGetStackArtifactParameters(artifactId *string) error {
	return nil
}

func (c *jsiiProxy_CloudAssembly) validateGetStackByNameParameters(stackName *string) error {
	return nil
}

func (c *jsiiProxy_CloudAssembly) validateTryGetArtifactParameters(id *string) error {
	return nil
}

func validateNewCloudAssemblyParameters(directory *string, loadOptions *cloudassemblyschema.LoadManifestOptions) error {
	return nil
}

