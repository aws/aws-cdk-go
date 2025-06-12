//go:build no_runtime_type_checking

package cxapi

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssetManifestArtifact) validateFindMetadataByTypeParameters(type_ *string) error {
	return nil
}

func validateAssetManifestArtifact_FromManifestParameters(assembly CloudAssembly, id *string, artifact *cloudassemblyschema.ArtifactManifest) error {
	return nil
}

func validateAssetManifestArtifact_IsAssetManifestArtifactParameters(art interface{}) error {
	return nil
}

func validateNewAssetManifestArtifactParameters(assembly CloudAssembly, name *string, artifact *cloudassemblyschema.ArtifactManifest) error {
	return nil
}

