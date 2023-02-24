package awsimagebuilder


// Define and configure faster launching for output Windows AMIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fastLaunchConfigurationProperty := &FastLaunchConfigurationProperty{
//   	AccountId: jsii.String("accountId"),
//   	Enabled: jsii.Boolean(false),
//   	LaunchTemplate: &FastLaunchLaunchTemplateSpecificationProperty{
//   		LaunchTemplateId: jsii.String("launchTemplateId"),
//   		LaunchTemplateName: jsii.String("launchTemplateName"),
//   		LaunchTemplateVersion: jsii.String("launchTemplateVersion"),
//   	},
//   	MaxParallelLaunches: jsii.Number(123),
//   	SnapshotConfiguration: &FastLaunchSnapshotConfigurationProperty{
//   		TargetResourceCount: jsii.Number(123),
//   	},
//   }
//
type CfnDistributionConfiguration_FastLaunchConfigurationProperty struct {
	// The owner account ID for the fast-launch enabled Windows AMI.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// A Boolean that represents the current state of faster launching for the Windows AMI.
	//
	// Set to `true` to start using Windows faster launching, or `false` to stop using it.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The launch template that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots.
	LaunchTemplate interface{} `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The maximum number of parallel instances that are launched for creating resources.
	MaxParallelLaunches *float64 `field:"optional" json:"maxParallelLaunches" yaml:"maxParallelLaunches"`
	// Configuration settings for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled.
	SnapshotConfiguration interface{} `field:"optional" json:"snapshotConfiguration" yaml:"snapshotConfiguration"`
}

