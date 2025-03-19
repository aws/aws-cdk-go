package awsimagebuilder


// Properties for defining a `CfnDistributionConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var amiDistributionConfiguration interface{}
//   var containerDistributionConfiguration interface{}
//   var ssmParameterConfigurations interface{}
//
//   cfnDistributionConfigurationProps := &CfnDistributionConfigurationProps{
//   	Distributions: []interface{}{
//   		&DistributionProperty{
//   			Region: jsii.String("region"),
//
//   			// the properties below are optional
//   			AmiDistributionConfiguration: amiDistributionConfiguration,
//   			ContainerDistributionConfiguration: containerDistributionConfiguration,
//   			FastLaunchConfigurations: []interface{}{
//   				&FastLaunchConfigurationProperty{
//   					AccountId: jsii.String("accountId"),
//   					Enabled: jsii.Boolean(false),
//   					LaunchTemplate: &FastLaunchLaunchTemplateSpecificationProperty{
//   						LaunchTemplateId: jsii.String("launchTemplateId"),
//   						LaunchTemplateName: jsii.String("launchTemplateName"),
//   						LaunchTemplateVersion: jsii.String("launchTemplateVersion"),
//   					},
//   					MaxParallelLaunches: jsii.Number(123),
//   					SnapshotConfiguration: &FastLaunchSnapshotConfigurationProperty{
//   						TargetResourceCount: jsii.Number(123),
//   					},
//   				},
//   			},
//   			LaunchTemplateConfigurations: []interface{}{
//   				&LaunchTemplateConfigurationProperty{
//   					AccountId: jsii.String("accountId"),
//   					LaunchTemplateId: jsii.String("launchTemplateId"),
//   					SetDefaultVersion: jsii.Boolean(false),
//   				},
//   			},
//   			LicenseConfigurationArns: []*string{
//   				jsii.String("licenseConfigurationArns"),
//   			},
//   			SsmParameterConfigurations: []interface{}{
//   				ssmParameterConfigurations,
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-distributionconfiguration.html
//
type CfnDistributionConfigurationProps struct {
	// The distributions of this distribution configuration formatted as an array of Distribution objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-distributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-distributions
	//
	Distributions interface{} `field:"required" json:"distributions" yaml:"distributions"`
	// The name of this distribution configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-distributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of this distribution configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-distributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags of this distribution configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-distributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

