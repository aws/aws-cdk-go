//go:build no_runtime_type_checking

package cxapi

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFormationStackArtifact) validateFindMetadataByTypeParameters(type_ *string) error {
	return nil
}

func validateCloudFormationStackArtifact_FromManifestParameters(assembly CloudAssembly, id *string, artifact *cloudassemblyschema.ArtifactManifest) error {
	return nil
}

func validateNewCloudFormationStackArtifactParameters(assembly CloudAssembly, artifactId *string, artifact *cloudassemblyschema.ArtifactManifest) error {
	return nil
}

