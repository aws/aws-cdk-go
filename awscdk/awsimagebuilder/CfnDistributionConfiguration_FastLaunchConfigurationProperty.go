package awsimagebuilder


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
	// `CfnDistributionConfiguration.FastLaunchConfigurationProperty.AccountId`.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// `CfnDistributionConfiguration.FastLaunchConfigurationProperty.Enabled`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// `CfnDistributionConfiguration.FastLaunchConfigurationProperty.LaunchTemplate`.
	LaunchTemplate interface{} `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// `CfnDistributionConfiguration.FastLaunchConfigurationProperty.MaxParallelLaunches`.
	MaxParallelLaunches *float64 `field:"optional" json:"maxParallelLaunches" yaml:"maxParallelLaunches"`
	// `CfnDistributionConfiguration.FastLaunchConfigurationProperty.SnapshotConfiguration`.
	SnapshotConfiguration interface{} `field:"optional" json:"snapshotConfiguration" yaml:"snapshotConfiguration"`
}

