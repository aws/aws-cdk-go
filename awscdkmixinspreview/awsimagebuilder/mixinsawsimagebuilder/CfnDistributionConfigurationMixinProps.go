package mixinsawsimagebuilder


// Properties for CfnDistributionConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var amiDistributionConfiguration interface{}
//   var containerDistributionConfiguration interface{}
//
//   cfnDistributionConfigurationMixinProps := &CfnDistributionConfigurationMixinProps{
//   	Description: jsii.String("description"),
//   	Distributions: []interface{}{
//   		&DistributionProperty{
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
//   			Region: jsii.String("region"),
//   			SsmParameterConfigurations: []interface{}{
//   				&SsmParameterConfigurationProperty{
//   					AmiAccountId: jsii.String("amiAccountId"),
//   					DataType: jsii.String("dataType"),
//   					ParameterName: jsii.String("parameterName"),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-distributionconfiguration.html
//
type CfnDistributionConfigurationMixinProps struct {
	// The description of this distribution configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-distributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The distributions of this distribution configuration formatted as an array of Distribution objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-distributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-distributions
	//
	Distributions interface{} `field:"optional" json:"distributions" yaml:"distributions"`
	// The name of this distribution configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-distributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags of this distribution configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-distributionconfiguration.html#cfn-imagebuilder-distributionconfiguration-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

