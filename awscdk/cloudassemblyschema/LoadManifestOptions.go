package cloudassemblyschema


// Options for the loadManifest operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadManifestOptions := &LoadManifestOptions{
//   	SkipEnumCheck: jsii.Boolean(false),
//   	SkipVersionCheck: jsii.Boolean(false),
//   	TopoSort: jsii.Boolean(false),
//   }
//
type LoadManifestOptions struct {
	// Skip enum checks.
	//
	// This means you may read enum values you don't know about yet. Make sure to always
	// check the values of enums you encounter in the manifest.
	// Default: false.
	//
	SkipEnumCheck *bool `field:"optional" json:"skipEnumCheck" yaml:"skipEnumCheck"`
	// Skip the version check.
	//
	// This means you may read a newer cloud assembly than the CX API is designed
	// to support, and your application may not be aware of all features that in use
	// in the Cloud Assembly.
	// Default: false.
	//
	SkipVersionCheck *bool `field:"optional" json:"skipVersionCheck" yaml:"skipVersionCheck"`
	// Topologically sort all artifacts.
	//
	// This parameter is only respected by the constructor of `CloudAssembly`. The
	// property lives here for backwards compatibility reasons.
	// Default: true.
	//
	TopoSort *bool `field:"optional" json:"topoSort" yaml:"topoSort"`
}

