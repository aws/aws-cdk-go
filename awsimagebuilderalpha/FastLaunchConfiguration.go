package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The EC2 Fast Launch configuration to use for the Windows AMI.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var launchTemplate LaunchTemplate
//
//   fastLaunchConfiguration := &FastLaunchConfiguration{
//   	Enabled: jsii.Boolean(false),
//   	LaunchTemplate: launchTemplate,
//   	MaxParallelLaunches: jsii.Number(123),
//   	TargetSnapshotCount: jsii.Number(123),
//   }
//
// Experimental.
type FastLaunchConfiguration struct {
	// Whether to enable fast launch for the AMI.
	// Default: false.
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The launch template that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots.
	// Default: None.
	//
	// Experimental.
	LaunchTemplate awsec2.ILaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The maximum number of parallel instances that are launched for creating resources.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EnableFastLaunch.html
	//
	// Default: A maximum of 6 instances are launched in parallel.
	//
	// Experimental.
	MaxParallelLaunches *float64 `field:"optional" json:"maxParallelLaunches" yaml:"maxParallelLaunches"`
	// The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EnableFastLaunch.html
	//
	// Default: 10 snapshots are kept pre-provisioned.
	//
	// Experimental.
	TargetSnapshotCount *float64 `field:"optional" json:"targetSnapshotCount" yaml:"targetSnapshotCount"`
}

