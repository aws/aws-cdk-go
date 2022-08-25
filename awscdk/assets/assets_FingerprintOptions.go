package assets

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Options related to calculating source hash.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fingerprintOptions := &fingerprintOptions{
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	extraHash: jsii.String("extraHash"),
//   	follow: awscdk.Assets.followMode_NEVER,
//   	ignoreMode: monocdk.ignoreMode_GLOB,
//   }
//
// Deprecated: see `core.FingerprintOptions`
type FingerprintOptions struct {
	// Glob patterns to exclude from the copy.
	// Deprecated: see `core.FingerprintOptions`
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	// Deprecated: use `followSymlinks` instead.
	Follow FollowMode `field:"optional" json:"follow" yaml:"follow"`
	// The ignore behavior to use for exclude patterns.
	// Deprecated: see `core.FingerprintOptions`
	IgnoreMode awscdk.IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	// Deprecated: see `core.FingerprintOptions`
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
}

