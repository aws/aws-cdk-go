// An experiment to bundle the entire CDK into a single module
package awscdk


// Options applied when copying directories into the staging location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   fileCopyOptions := &fileCopyOptions{
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	followSymlinks: monocdk.symlinkFollowMode_NEVER,
//   	ignoreMode: monocdk.ignoreMode_GLOB,
//   }
//
// Experimental.
type FileCopyOptions struct {
	// Glob patterns to exclude from the copy.
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	// Experimental.
	FollowSymlinks SymlinkFollowMode `field:"optional" json:"followSymlinks" yaml:"followSymlinks"`
	// The ignore behavior to use for exclude patterns.
	// Experimental.
	IgnoreMode IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
}

