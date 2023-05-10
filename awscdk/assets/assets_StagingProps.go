package assets

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Deprecated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stagingProps := &StagingProps{
//   	SourcePath: jsii.String("sourcePath"),
//
//   	// the properties below are optional
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	ExtraHash: jsii.String("extraHash"),
//   	Follow: awscdk.Assets.FollowMode_NEVER,
//   	IgnoreMode: monocdk.IgnoreMode_GLOB,
//   }
//
// Deprecated: use `core.AssetStagingProps`
type StagingProps struct {
	// Glob patterns to exclude from the copy.
	// Deprecated: use `core.AssetStagingProps`
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	// Deprecated: use `followSymlinks` instead.
	Follow FollowMode `field:"optional" json:"follow" yaml:"follow"`
	// The ignore behavior to use for exclude patterns.
	// Deprecated: use `core.AssetStagingProps`
	IgnoreMode awscdk.IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	// Deprecated: use `core.AssetStagingProps`
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
	// Local file or directory to stage.
	// Deprecated: use `core.AssetStagingProps`
	SourcePath *string `field:"required" json:"sourcePath" yaml:"sourcePath"`
}

