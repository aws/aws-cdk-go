package awscdk


// Options related to calculating source hash.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fingerprintOptions := &FingerprintOptions{
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	ExtraHash: jsii.String("extraHash"),
//   	Follow: cdk.SymlinkFollowMode_NEVER,
//   	IgnoreMode: cdk.IgnoreMode_GLOB,
//   }
//
type FingerprintOptions struct {
	// File paths matching the patterns will be excluded.
	//
	// See `ignoreMode` to set the matching behavior.
	// Has no effect on Assets bundled using the `bundling` property.
	// Default: - nothing is excluded.
	//
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	// Default: SymlinkFollowMode.NEVER
	//
	Follow SymlinkFollowMode `field:"optional" json:"follow" yaml:"follow"`
	// The ignore behavior to use for `exclude` patterns.
	// Default: IgnoreMode.GLOB
	//
	IgnoreMode IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	// Default: - hash is only based on source content.
	//
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
}

