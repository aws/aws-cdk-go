package mixinsawsimagebuilder


// The distribution configuration distribution defines the settings for a specific Region in the Distribution Configuration.
//
// You must specify whether the distribution is for an AMI or a container image. To do so, include exactly one of the following data types for your distribution:
//
// - amiDistributionConfiguration
// - containerDistributionConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var amiDistributionConfiguration interface{}
//   var containerDistributionConfiguration interface{}
//
//   distributionProperty := &DistributionProperty{
//   	AmiDistributionConfiguration: amiDistributionConfiguration,
//   	ContainerDistributionConfiguration: containerDistributionConfiguration,
//   	FastLaunchConfigurations: []interface{}{
//   		&FastLaunchConfigurationProperty{
//   			AccountId: jsii.String("accountId"),
//   			Enabled: jsii.Boolean(false),
//   			LaunchTemplate: &FastLaunchLaunchTemplateSpecificationProperty{
//   				LaunchTemplateId: jsii.String("launchTemplateId"),
//   				LaunchTemplateName: jsii.String("launchTemplateName"),
//   				LaunchTemplateVersion: jsii.String("launchTemplateVersion"),
//   			},
//   			MaxParallelLaunches: jsii.Number(123),
//   			SnapshotConfiguration: &FastLaunchSnapshotConfigurationProperty{
//   				TargetResourceCount: jsii.Number(123),
//   			},
//   		},
//   	},
//   	LaunchTemplateConfigurations: []interface{}{
//   		&LaunchTemplateConfigurationProperty{
//   			AccountId: jsii.String("accountId"),
//   			LaunchTemplateId: jsii.String("launchTemplateId"),
//   			SetDefaultVersion: jsii.Boolean(false),
//   		},
//   	},
//   	LicenseConfigurationArns: []*string{
//   		jsii.String("licenseConfigurationArns"),
//   	},
//   	Region: jsii.String("region"),
//   	SsmParameterConfigurations: []interface{}{
//   		&SsmParameterConfigurationProperty{
//   			AmiAccountId: jsii.String("amiAccountId"),
//   			DataType: jsii.String("dataType"),
//   			ParameterName: jsii.String("parameterName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html
//
type CfnDistributionConfigurationPropsMixin_DistributionProperty struct {
	// The specific AMI settings, such as launch permissions and AMI tags.
	//
	// For details, see example schema below.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-amidistributionconfiguration
	//
	AmiDistributionConfiguration interface{} `field:"optional" json:"amiDistributionConfiguration" yaml:"amiDistributionConfiguration"`
	// Container distribution settings for encryption, licensing, and sharing in a specific Region.
	//
	// For details, see example schema below.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-containerdistributionconfiguration
	//
	ContainerDistributionConfiguration interface{} `field:"optional" json:"containerDistributionConfiguration" yaml:"containerDistributionConfiguration"`
	// The Windows faster-launching configurations to use for AMI distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-fastlaunchconfigurations
	//
	FastLaunchConfigurations interface{} `field:"optional" json:"fastLaunchConfigurations" yaml:"fastLaunchConfigurations"`
	// A group of launchTemplateConfiguration settings that apply to image distribution for specified accounts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-launchtemplateconfigurations
	//
	LaunchTemplateConfigurations interface{} `field:"optional" json:"launchTemplateConfigurations" yaml:"launchTemplateConfigurations"`
	// The License Manager Configuration to associate with the AMI in the specified Region.
	//
	// For more information, see the [LicenseConfiguration API](https://docs.aws.amazon.com/license-manager/latest/APIReference/API_LicenseConfiguration.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-licenseconfigurationarns
	//
	LicenseConfigurationArns *[]*string `field:"optional" json:"licenseConfigurationArns" yaml:"licenseConfigurationArns"`
	// The target Region for the Distribution Configuration.
	//
	// For example, `eu-west-1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Contains settings to update AWS Systems Manager (SSM) Parameter Store Parameters with output AMI IDs from the build by target Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-distribution.html#cfn-imagebuilder-distributionconfiguration-distribution-ssmparameterconfigurations
	//
	SsmParameterConfigurations interface{} `field:"optional" json:"ssmParameterConfigurations" yaml:"ssmParameterConfigurations"`
}

