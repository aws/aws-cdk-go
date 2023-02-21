//go:build no_runtime_type_checking

package cloudassemblyschema

// Building without runtime type checking enabled, so all the below just return nil

func validateManifest_LoadAssemblyManifestParameters(filePath *string, options *LoadManifestOptions) error {
	return nil
}

func validateManifest_LoadAssetManifestParameters(filePath *string) error {
	return nil
}

func validateManifest_LoadIntegManifestParameters(filePath *string) error {
	return nil
}

func validateManifest_SaveAssemblyManifestParameters(manifest *AssemblyManifest, filePath *string) error {
	return nil
}

func validateManifest_SaveAssetManifestParameters(manifest *AssetManifest, filePath *string) error {
	return nil
}

func validateManifest_SaveIntegManifestParameters(manifest *IntegManifest, filePath *string) error {
	return nil
}

