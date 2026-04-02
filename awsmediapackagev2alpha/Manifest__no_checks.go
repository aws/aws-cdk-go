//go:build no_runtime_type_checking

package awsmediapackagev2alpha

// Building without runtime type checking enabled, so all the below just return nil

func validateManifest_DashParameters(manifest *DashManifestConfiguration) error {
	return nil
}

func validateManifest_HlsParameters(manifest *HlsManifestConfiguration) error {
	return nil
}

func validateManifest_LowLatencyHLSParameters(manifest *LowLatencyHlsManifestConfiguration) error {
	return nil
}

func validateManifest_MssParameters(manifest *MssManifestConfiguration) error {
	return nil
}

