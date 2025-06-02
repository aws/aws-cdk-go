package awscdk


// Options applied when copying directories into the staging location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileCopyOptions := &FileCopyOptions{
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	FollowSymlinks: cdk.SymlinkFollowMode_NEVER,
//   	IgnoreMode: cdk.IgnoreMode_GLOB,
//   }
//
type FileCopyOptions struct {
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
	FollowSymlinks SymlinkFollowMode `field:"optional" json:"followSymlinks" yaml:"followSymlinks"`
	// The ignore behavior to use for `exclude` patterns.
	// Default: IgnoreMode.GLOB
	//
	IgnoreMode IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
}

