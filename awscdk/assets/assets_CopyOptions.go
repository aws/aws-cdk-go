package assets

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Obtains applied when copying directories into the staging location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   copyOptions := &copyOptions{
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	follow: awscdk.Assets.followMode_NEVER,
//   	ignoreMode: monocdk.ignoreMode_GLOB,
//   }
//
// Deprecated: see `core.CopyOptions`
type CopyOptions struct {
	// Glob patterns to exclude from the copy.
	// Deprecated: see `core.CopyOptions`
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	// Deprecated: use `followSymlinks` instead.
	Follow FollowMode `field:"optional" json:"follow" yaml:"follow"`
	// The ignore behavior to use for exclude patterns.
	// Deprecated: see `core.CopyOptions`
	IgnoreMode awscdk.IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
}

