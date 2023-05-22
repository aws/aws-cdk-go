//go:build no_runtime_type_checking

package cxapi

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NestedCloudAssemblyArtifact) validateFindMetadataByTypeParameters(type_ *string) error {
	return nil
}

func validateNestedCloudAssemblyArtifact_FromManifestParameters(assembly CloudAssembly, id *string, artifact *cloudassemblyschema.ArtifactManifest) error {
	return nil
}

func validateNewNestedCloudAssemblyArtifactParameters(assembly CloudAssembly, name *string, artifact *cloudassemblyschema.ArtifactManifest) error {
	return nil
}

