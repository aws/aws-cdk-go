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
//
//   cfnDistributionConfigurationProps := &cfnDistributionConfigurationProps{
//   	distributions: []interface{}{
//   		&distributionProperty{
//   			region: jsii.String("region"),
//
//   			// the properties below are optional
//   			amiDistributionConfiguration: amiDistributionConfiguration,
//   			containerDistributionConfiguration: containerDistributionConfiguration,
//   			fastLaunchConfigurations: []interface{}{
//   				&fastLaunchConfigurationProperty{
//   					accountId: jsii.String("accountId"),
//   					enabled: jsii.Boolean(false),
//   					launchTemplate: &fastLaunchLaunchTemplateSpecificationProperty{
//   						launchTemplateId: jsii.String("launchTemplateId"),
//   						launchTemplateName: jsii.String("launchTemplateName"),
//   						launchTemplateVersion: jsii.String("launchTemplateVersion"),
//   					},
//   					maxParallelLaunches: jsii.Number(123),
//   					snapshotConfiguration: &fastLaunchSnapshotConfigurationProperty{
//   						targetResourceCount: jsii.Number(123),
//   					},
//   				},
//   			},
//   			launchTemplateConfigurations: []interface{}{
//   				&launchTemplateConfigurationProperty{
//   					accountId: jsii.String("accountId"),
//   					launchTemplateId: jsii.String("launchTemplateId"),
//   					setDefaultVersion: jsii.Boolean(false),
//   				},
//   			},
//   			licenseConfigurationArns: []*string{
//   				jsii.String("licenseConfigurationArns"),
//   			},
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnDistributionConfigurationProps struct {
	// The distributions of this distribution configuration formatted as an array of Distribution objects.
	Distributions interface{} `field:"required" json:"distributions" yaml:"distributions"`
	// The name of this distribution configuration.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of this distribution configuration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags of this distribution configuration.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

