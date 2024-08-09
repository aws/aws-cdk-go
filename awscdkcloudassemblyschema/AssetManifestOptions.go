package awscdkcloudassemblyschema


// Configuration options for the Asset Manifest.
type AssetManifestOptions struct {
	// SSM parameter where the bootstrap stack version number can be found.
	//
	// - If this value is not set, the bootstrap stack name must be known at
	//   deployment time so the stack version can be looked up from the stack
	//   outputs.
	// - If this value is set, the bootstrap stack can have any name because
	//   we won't need to look it up.
	// Default: - Bootstrap stack version number looked up.
	//
	BootstrapStackVersionSsmParameter *string `field:"optional" json:"bootstrapStackVersionSsmParameter" yaml:"bootstrapStackVersionSsmParameter"`
	// Version of bootstrap stack required to deploy this stack.
	// Default: - Version 1 (basic modern bootstrap stack).
	//
	RequiresBootstrapStackVersion *float64 `field:"optional" json:"requiresBootstrapStackVersion" yaml:"requiresBootstrapStackVersion"`
}

