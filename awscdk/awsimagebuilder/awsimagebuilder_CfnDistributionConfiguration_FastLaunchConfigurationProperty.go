package awsimagebuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fastLaunchConfigurationProperty := &fastLaunchConfigurationProperty{
//   	accountId: jsii.String("accountId"),
//   	enabled: jsii.Boolean(false),
//   	launchTemplate: &fastLaunchLaunchTemplateSpecificationProperty{
//   		launchTemplateId: jsii.String("launchTemplateId"),
//   		launchTemplateName: jsii.String("launchTemplateName"),
//   		launchTemplateVersion: jsii.String("launchTemplateVersion"),
//   	},
//   	maxParallelLaunches: jsii.Number(123),
//   	snapshotConfiguration: &fastLaunchSnapshotConfigurationProperty{
//   		targetResourceCount: jsii.Number(123),
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

