package cloudassemblyschema


// Artifact properties for the Asset Manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetManifestProperties := &AssetManifestProperties{
//   	File: jsii.String("file"),
//
//   	// the properties below are optional
//   	BootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   	RequiresBootstrapStackVersion: jsii.Number(123),
//   }
//
type AssetManifestProperties struct {
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
	// Filename of the asset manifest.
	File *string `field:"required" json:"file" yaml:"file"`
}

