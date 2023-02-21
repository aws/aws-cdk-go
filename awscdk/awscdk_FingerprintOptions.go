// An experiment to bundle the entire CDK into a single module
package awscdk


// Options related to calculating source hash.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   fingerprintOptions := &fingerprintOptions{
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	extraHash: jsii.String("extraHash"),
//   	follow: monocdk.symlinkFollowMode_NEVER,
//   	ignoreMode: monocdk.ignoreMode_GLOB,
//   }
//
// Experimental.
type FingerprintOptions struct {
	// Glob patterns to exclude from the copy.
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	// Experimental.
	Follow SymlinkFollowMode `field:"optional" json:"follow" yaml:"follow"`
	// The ignore behavior to use for exclude patterns.
	// Experimental.
	IgnoreMode IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	// Experimental.
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
}

